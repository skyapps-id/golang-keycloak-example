package main

import (
	"golang-keycloak/controller"
	"golang-keycloak/middleware"
	"golang-keycloak/pkg"
	"golang-keycloak/repository"
	"golang-keycloak/service"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Pkg
	keycloak := pkg.NewKeycloak()
	oidc := pkg.NewOpenID()

	// User
	userService := service.NewUserService(keycloak)
	userController := controller.NewUserController(userService)

	// Todo
	todoRepository := repository.NewTodoRepostiroy()
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	// Middlerware
	middlerware := middleware.NewMiddlewares(keycloak, oidc)

	// Validation
	validation := pkg.NewCustomValidator()

	e := echo.New()
	e.Validator = validation
	e.POST("/user/admin/login", userController.LoginAdmin)
	e.POST("/user/create", userController.Create)
	e.POST("/user/login", userController.Login)
	e.POST("/user/refresh-token", userController.RefreshToken)
	e.POST("/user/logout", userController.Logout)
	e.POST("/user/info", userController.Info)
	e.GET("/todo", todoController.Todo, middlerware.Authenticate)
	e.GET("/todo/with-oidc", todoController.Todo, middlerware.AuthenticateOIDC)

	e.Logger.Fatal(e.Start(":3000"))
}
