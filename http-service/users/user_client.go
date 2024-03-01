package users

import (
	"context"
	"fmt"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/http-service/config"
	"google.golang.org/grpc"
)

type (
	UserCreds  = any
	UserClient interface {
		Save(user *users_rpc.StoreUserRequest) (*users_rpc.StoreUserReply, error)
		FindByUsername(name string) (*users_rpc.StoreUserReply, error)
		FindUserByCreds(creds *UserCreds) (*users_rpc.StoreUserReply, error)
	}
)

type UserRpcClient struct {
	client users_rpc.UserMangmentClient
}

func NewUserHTTPClient() *UserRpcClient {
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

func (uc *UserRpcClient) FindByUsername(name string) (*users_rpc.StoreUserReply, error) {
	return nil, nil
}

func (uc *UserRpcClient) FindUserByCreds(creds *UserCreds) (*users_rpc.StoreUserReply, error) {
	return nil, nil
}
