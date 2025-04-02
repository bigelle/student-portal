package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/bigelle/student-portal/proto/auth"
	"github.com/joho/godotenv"
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
	env := flag.String("env", "", "set dot env for this program")
	flag.Parse()
	log.Println(*env)
	err := godotenv.Load(*env)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", os.Getenv("AUTH_SERVICE_ADDR"))
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
