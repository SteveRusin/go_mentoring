package users

import (
	"context"
	"fmt"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/http-service/config"
	"google.golang.org/grpc"
)

type (
	UserClient interface {
		Save(user *users_rpc.StoreUserRequest) (*users_rpc.StoreUserReply, error)
		FindUserByCreds(creds *users_rpc.GetUserRequest) (*users_rpc.GetUserReply, error)
		NewUploadImageClient(ctx context.Context, opts ...grpc.CallOption) (users_rpc.UserMangment_UploadImageClient, error)
	}
)

type UserRpcClient struct {
	client users_rpc.UserMangmentClient
}

func NewUserRpcClient() *UserRpcClient {
	config := config.GetUserServerConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.Host, config.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &UserRpcClient{
		client: users_rpc.NewUserMangmentClient(conn),
	}
}

func (uc *UserRpcClient) Save(user *users_rpc.StoreUserRequest) (*users_rpc.StoreUserReply, error) {
	res, err := uc.client.StoreUser(context.TODO(), user)
	return res, err
}

func (uc *UserRpcClient) FindUserByCreds(creds *users_rpc.GetUserRequest) (*users_rpc.GetUserReply, error) {
	res, err := uc.client.GetUser(context.TODO(), creds)
	return res, err
}

func (uc *UserRpcClient) NewUploadImageClient(ctx context.Context, opts ...grpc.CallOption) (users_rpc.UserMangment_UploadImageClient, error) {
	return uc.client.UploadImage(ctx, opts...)
}
