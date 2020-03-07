package ratelimit

import (
	"context"
	"math"
	"time"

	"github.com/aws/aws-sdk-go-v2/internal/timeconv"
)

const (
	defaultAttemptCost         = 1
	defaultLimiterSmooth       = 0.8
	defaultTimeBucketThreshold = time.Millisecond * 500
)

// AdaptiveAttemptRateLimit is an rate limiter that will introduce adaptive limited
// tokens based on throttling responses from a operation attempt result.
type AdaptiveAttemptRateLimit struct {
	// Rate at which tokens are replenished into the bucket.
	fillRate float64
	// Maximum tokens in bucket.
	maxCapacity float64

	// if the bucket rate limiting is enabled.
	enabled bool

	// Current capacity of the rate limited bucket.
	capacity float64
	// Last time the bucket was refilled.
	lastRefill time.Time

	// ??? Smoothed rate tokens are being retrieved.
	measuredTXRate float64
	// ??? Last half second bucket was used.
	lastTXRateBucket time.Time
	// number of requests since last measured rate update.
	requestCount float64

	// max rate when a most recent throttle occurred.
	lastThrottleMaxRate float64
	// when a throttle last occurred.
	lastThrottle time.Time

	requestCost float64

	timeNow             func() time.Time
	sleepFn             func(context.Context, time.Duration) error
	smooth              float64
	timeBucketThreshold time.Duration
}

func timeNow() time.Time {
	return time.Now()
}

// TODO move this into the SDK's core internal pkg.
func sleepFn(ctx context.Context, dur time.Duration) error {
	ticker := time.NewTimer(dur)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ticker.C:
		return nil
	}
}

// NewAdaptiveAttemptRateLimit returns a rate limiter for limiting operation
// attempt rates based the result of previous attempt invocation.
func NewAdaptiveAttemptRateLimit() *AdaptiveAttemptRateLimit {
	now := timeNow()

	return &AdaptiveAttemptRateLimit{
		lastTXRateBucket: now.Truncate(time.Second), // floor to seconds
		lastThrottle:     timeNow(),

		timeNow: timeNow,
		sleepFn: sleepFn,

		requestCost:         defaultAttemptCost,
		smooth:              defaultLimiterSmooth,
		timeBucketThreshold: defaultTimeBucketThreshold,
	}
}

// GetAttemptToken will block until a token is available, or until context is
// canceled. Will return an error if canceled before a token is available.
func (l *AdaptiveAttemptRateLimit) GetAttemptToken(ctx context.Context) (func(error) error, error) {
	delay := l.getTokenDelay(l.requestCost)
	if delay == 0 {
		return l.releaseAttemptToken, nil
	}

	if err := l.sleepFn(ctx, delay); err != nil {
		return nil, err
	}

	return l.releaseAttemptToken, nil
}

// Computes the amount of delay based on the amount of tokens requested.
// Returns 0 if not enabled, or there are enough tokens available.
func (l *AdaptiveAttemptRateLimit) getTokenDelay(amount float64) time.Duration {
	if !l.enabled {
		return 0
	}

	var delay time.Duration

	l.refillBucket()
	if amount > l.capacity {
		delay = timeconv.FloatSecondsDur((amount - l.capacity) / l.fillRate)
	}

	l.capacity -= amount
	return delay
}

// refills the bucket based on the current capacity and time since last time
// refill has occurred.
func (l *AdaptiveAttemptRateLimit) refillBucket() {
	now := l.timeNow()

	if !l.lastRefill.IsZero() {
		fillAmount := timeconv.DurSecondsFloat(now.Sub(l.lastRefill)) * l.fillRate
		l.capacity = math.Min(l.maxCapacity, l.capacity+fillAmount)
	}

	l.lastRefill = now
}

// ReleaseAttemptToken will return a token back to the pool, and adjust the rate
// future tokens can be retrieved based on the kind of result, and error that
// occurred while invoking the operation attempt.
func (l *AdaptiveAttemptRateLimit) releaseAttemptToken(err error) error {
	// TODO _UpdateClientSendingRate
	return nil
}

// Updates rate
func (l *AdaptiveAttemptRateLimit) updateMeasuredRate() {
	timeBucket := l.timeNow().Round(l.timeBucketThreshold)
	l.requestCount++

	if timeBucket.After(l.lastTXRateBucket) {
		currentRate := l.requestCount / timeconv.DurSecondsFloat(timeBucket.Sub(l.lastTXRateBucket))
		l.measuredTXRate = (currentRate * l.smooth) + (l.measuredTXRate * (1 - l.smooth))
		l.requestCount = 0
		l.lastTXRateBucket = timeBucket
	}
}

func (l *AdaptiveAttemptRateLimit) updateTokenBucketRate(newRPS float64) {
	// TODO _TokenBucketUpdateRate
}
