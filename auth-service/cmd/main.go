package main

import (
	"context"
	"log"
	"net"

	"github.com/bigelle/student-portal/proto/auth"
	"google.golang.org/grpc"
)

type authService struct {
	auth.UnimplementedAuthServer
}

func (s authService) Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{
		Name:    "it works",
		UserId:  42,
		Session: "session",
		Csrf:    "csrf",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := &authService{}

	auth.RegisterAuthServer(server, service)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
