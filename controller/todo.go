package controller

import (
	"golang-keycloak/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	TodoController struct {
		services service.TodoSerivce
	}
)

func NewTodoController(services service.TodoSerivce) *TodoController {
	return &TodoController{services: services}
}

func (c *TodoController) Todo(ctx echo.Context) error {
	result, err := c.services.Fatch(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}
