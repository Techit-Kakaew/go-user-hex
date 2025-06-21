package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/yourusername/go-user-hex/internal/user/handler"
	"github.com/yourusername/go-user-hex/internal/user/repository"
	"github.com/yourusername/go-user-hex/internal/user/usecase"
)

func main() {
	e := echo.New()

	// DI
	userRepo := repository.NewUserRepoMemory()
	userUC := usecase.NewUserUseCase(userRepo)
	handler.NewUserHandler(e, userUC)

	log.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
