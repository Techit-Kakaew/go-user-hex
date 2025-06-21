package handler

import (
	"net/http"

	"github.com/Techit-Kakaew/go-user-hex/internal/auth"
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
	e.POST("/api/users/login", handler.Login)
	e.GET("/api/users/me", handler.Me, auth.JWTMiddleware)
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

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(loginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	token, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *UserHandler) Me(c echo.Context) error {
	userID := c.Get("userID").(string)

	user, err := h.usecase.GetByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
