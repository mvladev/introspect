// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package talent

import (
	"context"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// EventCallOptions contains the retry settings for each method of EventClient.
type EventCallOptions struct {
	CreateClientEvent []gax.CallOption
}

func defaultEventClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("jobs.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultEventCallOptions() *EventCallOptions {
	retry := map[[2]string][]gax.CallOption{}
	return &EventCallOptions{
		CreateClientEvent: retry[[2]string{"default", "non_idempotent"}],
	}
}

// EventClient is a client for interacting with Cloud Talent Solution API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type EventClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	eventClient talentpb.EventServiceClient

	// The call options for this service.
	CallOptions *EventCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewEventClient creates a new event service client.
//
// A service handles client event report.
func NewEventClient(ctx context.Context, opts ...option.ClientOption) (*EventClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultEventClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &EventClient{
		conn:        conn,
		CallOptions: defaultEventCallOptions(),

		eventClient: talentpb.NewEventServiceClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *EventClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *EventClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *EventClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateClientEvent report events issued when end user interacts with customer's application
// that uses Cloud Talent Solution. You may inspect the created events in
// self service
// tools (at https://console.cloud.google.com/talent-solution/overview).
// Learn
// more (at https://cloud.google.com/talent-solution/job-search/docs/management-tools)
// about self service tools.
func (c *EventClient) CreateClientEvent(ctx context.Context, req *talentpb.CreateClientEventRequest, opts ...gax.CallOption) (*talentpb.ClientEvent, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateClientEvent[0:len(c.CallOptions.CreateClientEvent):len(c.CallOptions.CreateClientEvent)], opts...)
	var resp *talentpb.ClientEvent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.eventClient.CreateClientEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
