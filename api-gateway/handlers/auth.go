package handlers

import (
	"net/http"

	"github.com/bigelle/student-portal/api-gateway/schemas"
	"github.com/bigelle/student-portal/proto/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthClient auth.AuthClient
}

func (h *AuthHandler) HandleLogin(c echo.Context) error {
	var l schemas.LoginRequest
	err := c.Bind(&l)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := l.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	req := auth.LoginRequest{
		Name:     l.Name,
		Password: l.Password,
	}
	resp, err := h.AuthClient.Login(c.Request().Context(), &req)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schemas.LoginResponse{
		Name: resp.Name,
		ID:   resp.UserId,
		Role: resp.Role,
	})
}
