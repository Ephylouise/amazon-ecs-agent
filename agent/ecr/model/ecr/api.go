// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Code generated by [agent/gogenerate/awssdk.go] DO NOT EDIT.

package ecr

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
)

const opGetAuthorizationToken = "GetAuthorizationToken"

// GetAuthorizationTokenRequest generates a "aws/request.Request" representing the
// client's request for the GetAuthorizationToken operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetAuthorizationToken for more information on using the GetAuthorizationToken
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the GetAuthorizationTokenRequest method.
//	req, resp := client.GetAuthorizationTokenRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECR) GetAuthorizationTokenRequest(input *GetAuthorizationTokenInput) (req *request.Request, output *GetAuthorizationTokenOutput) {
	op := &request.Operation{
		Name:       opGetAuthorizationToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAuthorizationTokenInput{}
	}

	output = &GetAuthorizationTokenOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetAuthorizationToken API operation for Amazon EC2 Container Registry.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Registry's
// API operation GetAuthorizationToken for usage and error information.
//
// Returned Error Types:
//
//   - ServerException
//
//   - InvalidParameterException
func (c *ECR) GetAuthorizationToken(input *GetAuthorizationTokenInput) (*GetAuthorizationTokenOutput, error) {
	req, out := c.GetAuthorizationTokenRequest(input)
	return out, req.Send()
}

// GetAuthorizationTokenWithContext is the same as GetAuthorizationToken with the addition of
// the ability to pass a context and additional request options.
//
// See GetAuthorizationToken for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECR) GetAuthorizationTokenWithContext(ctx aws.Context, input *GetAuthorizationTokenInput, opts ...request.Option) (*GetAuthorizationTokenOutput, error) {
	req, out := c.GetAuthorizationTokenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AuthorizationData struct {
	_ struct{} `type:"structure"`

	AuthorizationToken *string `locationName:"authorizationToken" type:"string"`

	ExpiresAt *time.Time `locationName:"expiresAt" type:"timestamp"`

	ProxyEndpoint *string `locationName:"proxyEndpoint" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuthorizationData) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuthorizationData) GoString() string {
	return s.String()
}

// SetAuthorizationToken sets the AuthorizationToken field's value.
func (s *AuthorizationData) SetAuthorizationToken(v string) *AuthorizationData {
	s.AuthorizationToken = &v
	return s
}

// SetExpiresAt sets the ExpiresAt field's value.
func (s *AuthorizationData) SetExpiresAt(v time.Time) *AuthorizationData {
	s.ExpiresAt = &v
	return s
}

// SetProxyEndpoint sets the ProxyEndpoint field's value.
func (s *AuthorizationData) SetProxyEndpoint(v string) *AuthorizationData {
	s.ProxyEndpoint = &v
	return s
}

type GetAuthorizationTokenInput struct {
	_ struct{} `type:"structure"`

	RegistryIds []*string `locationName:"registryIds" type:"list"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetAuthorizationTokenInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetAuthorizationTokenInput) GoString() string {
	return s.String()
}

// SetRegistryIds sets the RegistryIds field's value.
func (s *GetAuthorizationTokenInput) SetRegistryIds(v []*string) *GetAuthorizationTokenInput {
	s.RegistryIds = v
	return s
}

type GetAuthorizationTokenOutput struct {
	_ struct{} `type:"structure"`

	AuthorizationData []*AuthorizationData `locationName:"authorizationData" type:"list"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetAuthorizationTokenOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetAuthorizationTokenOutput) GoString() string {
	return s.String()
}

// SetAuthorizationData sets the AuthorizationData field's value.
func (s *GetAuthorizationTokenOutput) SetAuthorizationData(v []*AuthorizationData) *GetAuthorizationTokenOutput {
	s.AuthorizationData = v
	return s
}

type InvalidParameterException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidParameterException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidParameterException) GoString() string {
	return s.String()
}

func newErrorInvalidParameterException(v protocol.ResponseMetadata) error {
	return &InvalidParameterException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidParameterException) Code() string {
	return "InvalidParameterException"
}

// Message returns the exception's message.
func (s *InvalidParameterException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidParameterException) OrigErr() error {
	return nil
}

func (s *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidParameterException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidParameterException) RequestID() string {
	return s.RespMetadata.RequestID
}

type ServerException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ServerException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ServerException) GoString() string {
	return s.String()
}

func newErrorServerException(v protocol.ResponseMetadata) error {
	return &ServerException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ServerException) Code() string {
	return "ServerException"
}

// Message returns the exception's message.
func (s *ServerException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ServerException) OrigErr() error {
	return nil
}

func (s *ServerException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ServerException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ServerException) RequestID() string {
	return s.RespMetadata.RequestID
}
