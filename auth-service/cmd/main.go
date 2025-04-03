package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/bigelle/student-portal/auth-service/handlers"
	"github.com/bigelle/student-portal/proto/auth"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

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
	service := &handlers.AuthService{}

	auth.RegisterAuthServer(server, service)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
