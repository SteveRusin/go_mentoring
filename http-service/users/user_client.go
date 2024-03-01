package users

import (
	"fmt"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/http-service/config"
	"google.golang.org/grpc"
)

type (
	User       = users_rpc.StoreUserRequest
	UserCreds  = any
	UserClient interface {
		Save(user User) (*User, error)
		FindByUsername(name string) (*User, error)
		FindUserByCreds(creds *UserCreds) (*User, error)
	}
)

type UserRpcClient struct {
	client users_rpc.UserMangmentClient
}

func NewUserHTTPClient() *UserRpcClient {
	config := config.GetUserServerConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		panic(err)
	}

	return &UserRpcClient{
		client: users_rpc.NewUserMangmentClient(conn),
	}
}

func (uc *UserRpcClient) Save(user User) (*User, error) {
	res, err := uc.Save(user)
	return res, err
}

func (uc *UserRpcClient) FindByUsername(name string) (*User, error) {
	return nil, nil
}

func (uc *UserRpcClient) FindUserByCreds(creds *UserCreds) (*User, error) {
	return nil, nil
}
