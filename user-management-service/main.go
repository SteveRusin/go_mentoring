package main

import (
	"fmt"
	"log"
	"net"

	"github.com/SteveRusin/go_mentoring/user-management-service/config"
	"github.com/SteveRusin/go_mentoring/user-management-service/repository"
	_ "github.com/joho/godotenv/autoload" // read .env file
	"google.golang.org/grpc"
	users_rpc "github.com/SteveRusin/go_mentoring/generated"
)

func main() {
	repository.MigrateUsersDb()
	appConfig := config.GetAppConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
  users_rpc.RegisterUserMangmentServer(grpcServer, newServer())

	log.Printf("Server is listening on %s:%s", appConfig.Host, appConfig.Port)
	grpcServer.Serve(lis)
}
