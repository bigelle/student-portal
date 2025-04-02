package main

import (
	"flag"
	"log"
	"os"

	"github.com/bigelle/student-portal/api-gateway/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	env := flag.String("env", "", "dot env for the executable")
	flag.Parse()
	if err := godotenv.Load(*env); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger())
	g := e.Group("/api")

	authConn, err := grpc.NewClient(os.Getenv("AUTH_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer authConn.Close()
	if err := server.SetupGroup(g, "/auth", authConn); err != nil {
		log.Fatal(err)
	}

	e.Start(":8080")
}
