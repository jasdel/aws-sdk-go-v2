// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateChangeset struct {
}

func (*awsRestjson1_serializeOpCreateChangeset) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateChangeset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateChangesetInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/datasets/{datasetId}/changesets")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsCreateChangesetInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateChangesetInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateChangesetInput(v *CreateChangesetInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.DatasetId == nil || len(*v.DatasetId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member datasetId must not be empty")}
	}
	if v.DatasetId != nil {
		if err := encoder.SetURI("datasetId").String(*v.DatasetId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateChangesetInput(v *CreateChangesetInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.ChangeType) > 0 {
		ok := object.Key("changeType")
		ok.String(string(v.ChangeType))
	}

	if v.FormatParams != nil {
		ok := object.Key("formatParams")
		if err := awsRestjson1_serializeDocumentStringMap(v.FormatParams, ok); err != nil {
			return err
		}
	}

	if len(v.FormatType) > 0 {
		ok := object.Key("formatType")
		ok.String(string(v.FormatType))
	}

	if v.SourceParams != nil {
		ok := object.Key("sourceParams")
		if err := awsRestjson1_serializeDocumentStringMap(v.SourceParams, ok); err != nil {
			return err
		}
	}

	if len(v.SourceType) > 0 {
		ok := object.Key("sourceType")
		ok.String(string(v.SourceType))
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentStringMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetProgrammaticAccessCredentials struct {
}

func (*awsRestjson1_serializeOpGetProgrammaticAccessCredentials) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetProgrammaticAccessCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetProgrammaticAccessCredentialsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/credentials/programmatic")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetProgrammaticAccessCredentialsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetProgrammaticAccessCredentialsInput(v *GetProgrammaticAccessCredentialsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.DurationInMinutes != 0 {
		encoder.SetQuery("durationInMinutes").Long(v.DurationInMinutes)
	}

	if v.EnvironmentId != nil {
		encoder.SetQuery("environmentId").String(*v.EnvironmentId)
	}

	return nil
}

type awsRestjson1_serializeOpGetWorkingLocation struct {
}

func (*awsRestjson1_serializeOpGetWorkingLocation) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetWorkingLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetWorkingLocationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/workingLocationV1")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetWorkingLocationInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetWorkingLocationInput(v *GetWorkingLocationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetWorkingLocationInput(v *GetWorkingLocationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.LocationType) > 0 {
		ok := object.Key("locationType")
		ok.String(string(v.LocationType))
	}

	return nil
}

func awsRestjson1_serializeDocumentStringMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}
