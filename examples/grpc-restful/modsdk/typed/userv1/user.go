// Code generated by protoc-gen-grpc-gateway-gosdk. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-gosdk 1.8.1-next-e99dba
// type: user

package userv1

import (
	"context"

	"github.com/jaronnie/autosdk/pb/userpb"
	"github.com/jaronnie/autosdk/rest"
)

type UserGetter interface {
	User() UserInterface
}

type UserInterface interface {
	// Add trans *userpb.AddUserReq into *userpb.AddUserResp
	// API /api/v1.0/user/add
	Add(ctx context.Context, param *userpb.AddUserReq) (*userpb.AddUserResp, error)

	UserExpansion
}

type userClient struct {
	client rest.Interface
}

func newUserClient(c *Userv1Client) *userClient {
	return &userClient{
		client: c.RESTClient(),
	}
}

func (x *userClient) Add(ctx context.Context, param *userpb.AddUserReq) (*userpb.AddUserResp, error) {
	var resp userpb.AddUserResp
	err := x.client.Verb("POST").
		SubPath(
			"/api/v1.0/user/add",
		).
		Params().
		Body(param).
		Do(ctx).
		Into(&resp, false)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}