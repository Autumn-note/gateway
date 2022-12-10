package client

import (
	"context"
	user "gateway/internal/grpc/pb"
	"gateway/pkg/grpc_helper"
	"google.golang.org/grpc"
)

const (
	serviceName = "0.0.0.0"
	port        = "50051"
	address     = serviceName + ":" + port
)

type UserClient struct {
	client user.UserServiceClient
	conn   *grpc.ClientConn
}

func NewClient() (*UserClient, error) {
	return NewClientWithAddress(address)
}

func NewClientWithAddress(address string) (*UserClient, error) {
	conn, err := grpc_helper.NewInsecureConnection(address)
	if err != nil {
		return nil, err
	}
	client := user.NewUserServiceClient(conn)
	return &UserClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *UserClient) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	return c.client.GetUser(ctx, req)
}
