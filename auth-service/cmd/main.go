package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type authService struct{}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
}
