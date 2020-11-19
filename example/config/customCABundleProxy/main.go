package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/awslabs/smithy-go/middleware"
)

func main() {
	cleanupCA, err := setupCAEnv()
	if err != nil {
		log.Fatalf("setup CA bundle env failed, %v", err)
	}
	defer cleanupCA()

	//os.Setenv("HTTP_PROXY", "http://127.0.0.1:8080")
	//os.Setenv("HTTPS_PROXY", "https://127.0.0.1:8080")

	if err := sdkRequest(); err != nil {
		log.Fatalf("sdk request failed, %v", err)
	}
}

func setupCAEnv() (func(), error) {
	cert, key, ca, err := awstesting.CreateTLSBundleFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to create CA bundle, %w", err)
	}

	return func() {
		awstesting.CleanupTLSBundleFiles(cert, key, ca)
	}, nil
}

func sdkRequest() error {
	cfg, err := config.LoadDefaultConfig(
		config.WithClientLogMode(
			aws.LogSigning|aws.LogRetries|aws.LogRequest|aws.LogResponseWithBody,
		),
		config.WithAPIOptions([]func(*middleware.Stack) error{
			func(stack *middleware.Stack) error {
				stack.Finalize.Remove("DisableAcceptEncodingGzip")
				return nil
			},
		}),
	)
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %w", err)
	}

	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to retrieve credentials, %w", err)
	}
	fmt.Printf("[INFO] Using AccessKeyID %s loaded from %s\n", creds.AccessKeyID, creds.Source)

	svc := s3.NewFromConfig(cfg)

	// Build the request with its input parameters
	resp, err := svc.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("failed to list buckets, %w", err)
	}

	fmt.Println("Buckets:", len(resp.Buckets))

	return nil
}
