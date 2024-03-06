package main

import (
	"context"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/user-management-service/randomId"
	"github.com/SteveRusin/go_mentoring/user-management-service/repository"
)

type Server struct {
	users_rpc.UnimplementedUserMangmentServer
	repo repository.UserRepository
}

func newServer() *Server {
	return &Server{
		repo: repository.NewUserPgRepository(),
	}
}

func (s *Server) StoreUser(ctx context.Context, req *users_rpc.StoreUserRequest) (*users_rpc.StoreUserReply, error) {
	user := &repository.User{
		Id:       randomId.New(),
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	result, err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}

	reply := users_rpc.StoreUserReply{
		Id:   result.Id,
		Name: result.Name,
	}

	return &reply, nil
}

func (s *Server) GetUser(ctx context.Context, req *users_rpc.GetUserRequest) (*users_rpc.GetUserReply, error) {
	creds := &repository.UserCreds{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	res, err := s.repo.FindUserByCreds(creds)
	if err != nil {
		return nil, err
	}

	reply := &users_rpc.GetUserReply{
		Id:   res.Id,
		Name: res.Name,
	}

	return reply, nil
}
