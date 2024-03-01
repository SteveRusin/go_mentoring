package main

import (
	"context"
	"log"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
)

type Server struct{
  users_rpc.UnimplementedUserMangmentServer
}


func newServer() *Server {
  return &Server{}
}

func (s *Server) StoreUser(ctx context.Context, req *users_rpc.StoreUserRequest) (*users_rpc.StoreUserReply, error) {
  log.Println("StoreUser called")
  return nil, nil
}
