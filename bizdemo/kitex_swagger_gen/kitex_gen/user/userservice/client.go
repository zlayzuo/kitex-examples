// Code generated by Kitex v0.12.0. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/cloudwego/kitex-examples/bizdemo/kitex_swagger_gen/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UpdateUser(ctx context.Context, req *user.UpdateUserRequest, callOptions ...callopt.Option) (r *user.UpdateUserResponse, err error)
	DeleteUser(ctx context.Context, req *user.DeleteUserRequest, callOptions ...callopt.Option) (r *user.DeleteUserResponse, err error)
	QueryUser(ctx context.Context, req *user.QueryUserRequest, callOptions ...callopt.Option) (r *user.QueryUserResponse, err error)
	CreateUser(ctx context.Context, req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, callOptions ...callopt.Option) (r *user.UpdateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, req)
}

func (p *kUserServiceClient) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, callOptions ...callopt.Option) (r *user.DeleteUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, req)
}

func (p *kUserServiceClient) QueryUser(ctx context.Context, req *user.QueryUserRequest, callOptions ...callopt.Option) (r *user.QueryUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUser(ctx, req)
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}
