package handlers

import (
	"context"
	"fmt"

	"github.com/bigelle/student-portal/proto/auth"
)

type AuthService struct {
	auth.UnimplementedAuthServer
}

func (s AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	l := loginRequest{Name: req.Name, Password: req.Password}
	if err := l.Validate(); err != nil {
		return nil, err
	}
	return &auth.LoginResponse{
		Name:    l.Name,
		UserId:  42,
		Role:    "master",
		Session: "session",
		Csrf:    "csrf",
	}, nil
}

type loginRequest struct {
	Name     string
	Password string
}

func (l loginRequest) Validate() error {
	if l.Name == "" {
		return fmt.Errorf("name can't be empty - saying from auth service")
	}
	if len(l.Password) < 8 {
		return fmt.Errorf("password must be 8 characters long")
	}
	return nil
}
