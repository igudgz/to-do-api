package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todo-api/auth"
	"todo-api/config"
	"todo-api/config/database"
	"todo-api/handler"
	"todo-api/persistence"
	"todo-api/service"
)

func main() {
	e := echo.New()
	auth.WebSecurityConfig(e)
	config := config.New()
	db := database.DBConnect(config)

	userPesistence := persistence.NewPersistenceUser(db, config)
	userService := service.NewUserService(userPesistence)
	userHandler := handler.NewUserHandler(userService)

	configureUserRoutes(e, userHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Port)))
}

func configureUserRoutes(e *echo.Echo, userHandler handler.UserHandler) {
	v1 := e.Group("/api/v1")
	v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())
	v1.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{Generator: func() string {
		return uuid.New().String()
	}}))

	v1.POST("/login", userHandler.AuthenticateUser)
}
