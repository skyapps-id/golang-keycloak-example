package controller

import (
	"golang-keycloak/dto"
	"golang-keycloak/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserController struct {
		services service.UserSerivce
	}
)

func NewUserController(services service.UserSerivce) *UserController {
	return &UserController{services: services}
}

func (c *UserController) LoginAdmin(ctx echo.Context) error {
	var req dto.ReqUserLogin
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := c.services.LoginAdmin(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *UserController) Create(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")

	result, err := c.services.Create(ctx.Request().Context(), token)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *UserController) Login(ctx echo.Context) error {
	var req dto.ReqUserLogin
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := c.services.Login(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *UserController) RefreshToken(ctx echo.Context) error {
	var req dto.ReqUserRefreshTokenOrLogout
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := c.services.RefreshToken(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *UserController) Logout(ctx echo.Context) error {
	var req dto.ReqUserRefreshTokenOrLogout
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := c.services.Logout(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *UserController) Info(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")

	result, err := c.services.Info(ctx.Request().Context(), token)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, result)
}
