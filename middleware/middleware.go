package middleware

import (
	"fmt"
	"golang-keycloak/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Middlewares interface {
		Authenticate(next echo.HandlerFunc) echo.HandlerFunc
	}

	middlewaresInst struct {
		keycloak pkg.Keycloak
	}
)

func NewMiddlewares(keycloak pkg.Keycloak) Middlewares {
	return &middlewaresInst{keycloak: keycloak}
}

func (mw *middlewaresInst) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		if token != "" {
			result, err := mw.keycloak.Gocloak.RetrospectToken(ctx.Request().Context(), token, mw.keycloak.ClientID, mw.keycloak.ClientSecret, mw.keycloak.Realm)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			jwt, _, err := mw.keycloak.Gocloak.DecodeAccessToken(ctx.Request().Context(), token, mw.keycloak.Realm)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			fmt.Println(jwt.Claims)

			if !*result.Active {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")

			}

			return next(ctx)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}
	}
}
