package client

import (
	"github.com/pkg/errors"
	demoservice "github.com/rantav/go-grpc-channelz/internal/generated/service"
	"google.golang.org/grpc"
)

// New creates a new gRPC client
func New(connectionString string) (demoservice.DemoServiceClient, error) {
	return NewWithDialOpts(connectionString)
}

// NewWithDialOpts creates a new gRPC client with custom []grpc.DialOption
func NewWithDialOpts(connectionString string, dialOpts ...grpc.DialOption) (demoservice.DemoServiceClient, error) {
	conn, err := grpc.Dial(connectionString, dialOpts...)
	if err != nil {
		return nil, errors.Wrapf(err, "error dialing to %s", connectionString)
	}

	client := demoservice.NewDemoServiceClient(conn)
	return client, nil
}
