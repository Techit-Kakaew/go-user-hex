package handler

import (
	"net/http"

	"github.com/Techit-Kakaew/go-user-hex/internal/user/domain"
	"github.com/Techit-Kakaew/go-user-hex/internal/user/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(e *echo.Echo, uc usecase.UserUseCase) {
	handler := &UserHandler{usecase: uc}

	e.POST("/api/users/register", handler.Register)
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(c echo.Context) error {
	req := new(registerRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.usecase.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "user registered"})
}
