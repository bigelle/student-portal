package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/bigelle/student-portal/api-gateway/schemas"
	"github.com/bigelle/student-portal/proto/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthClient auth.AuthClient
}

func (h *AuthHandler) HandleLogin(c echo.Context) error {
	var l schemas.Login
	err := c.Bind(&l)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	req := auth.LoginRequest{
		Name:     l.Name,
		Password: l.Password,
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()
	resp, err := h.AuthClient.Login(ctx, &req)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
