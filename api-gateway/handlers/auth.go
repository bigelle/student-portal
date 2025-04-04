package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bigelle/student-portal/api-gateway/schemas"
	"github.com/bigelle/student-portal/proto/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthClient auth.AuthClient
}

func (h *AuthHandler) HandleRegister(c echo.Context) error {
	var l schemas.RegisterRequest
	err := c.Bind(&l)
	if err != nil {
		return c.JSON(http.StatusOK, schemas.Response{
			Ok:          false,
			Description: "bad request",
		})
	}
	if err := l.Validate(); err != nil {
		return c.JSON(http.StatusOK, schemas.Response{
			Ok:          false,
			Description: fmt.Sprintf("bad request: %s", err.Error()),
		})
	}

	req := auth.RegisterRequest{}
	resp, err := h.sendRegisterReq(c.Request().Context(), &req)
	if err != nil || !resp.Ok {
		return c.JSON(http.StatusOK, schemas.Response{
			Ok:          false,
			Description: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, schemas.Response{
		Ok: true,
		Result: schemas.RegisterResponse{
			NewLogin: resp.NewLogin,
			UserID:   4269,
			Role:     "role",
		},
	})
}

func (h AuthHandler) sendRegisterReq(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	m := os.Getenv("USE_MOCKS")
	if m == "true" || m == "" {
		return fakeRegister(ctx, req)
	}
	return h.register(ctx, req)
}

func (h AuthHandler) register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return h.AuthClient.Register(ctx, req)
}

func fakeRegister(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		return &auth.RegisterResponse{}, nil
	}
}

func (h *AuthHandler) HandleLogin(c echo.Context) error {
	var l schemas.LoginRequest
	err := c.Bind(&l)
	if err != nil {
		return c.JSON(http.StatusOK, schemas.Response{
			Ok:          false,
			Description: "bad request",
		})
	}
	if err := l.Validate(); err != nil {
		return c.JSON(http.StatusOK, schemas.Response{
			Ok:          false,
			Description: fmt.Sprintf("bad request: %s", err.Error()),
		})
	}

	req := auth.LoginRequest{
		Name:     l.Name,
		Password: l.Password,
	}
	resp, err := h.sendLoginReq(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusOK, schemas.Response{
			Ok:          false,
			Description: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, schemas.Response{
		Ok: true,
		Result: schemas.LoginResponse{
			Name: resp.Name,
			ID:   resp.UserId,
			Role: resp.Role,
		},
	})
}

func (h AuthHandler) sendLoginReq(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	m := os.Getenv("USE_MOCKS")
	if m == "true" || m == "" {
		return fakeLogin(ctx, req)
	}
	return h.login(ctx, req)
}

func (h AuthHandler) login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return h.AuthClient.Login(ctx, req)
}

func fakeLogin(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		return &auth.LoginResponse{
			Name:    req.Name,
			UserId:  42,
			Role:    "mocker",
			Session: "session",
			Csrf:    "csrf",
		}, nil
	}
}
