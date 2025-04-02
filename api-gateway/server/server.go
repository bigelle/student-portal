package server

import (
	"fmt"

	"github.com/bigelle/student-portal/api-gateway/handlers"
	"github.com/bigelle/student-portal/proto/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func SetupGroup(g *echo.Group, group string, conn *grpc.ClientConn) error {
	switch group {
	case "/auth":
		setupAuth(g, conn)
		return nil

	default:
		return fmt.Errorf("unexpected group: %s", group)
	}
}

func setupAuth(g *echo.Group, conn *grpc.ClientConn) {
	authClient := auth.NewAuthClient(conn)
	h := handlers.AuthHandler{AuthClient: authClient}

	au := g.Group("/auth")
	au.GET("/login", h.HandleLogin)
}
