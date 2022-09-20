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
		AuthenticateOIDC(next echo.HandlerFunc) echo.HandlerFunc
	}

	middlewaresInst struct {
		keycloak pkg.Keycloak
		oidc     pkg.OpenID
	}

	clientRoles struct {
		Roles []string `json:"roles,omitempty"`
	}
	client struct {
		GolangApp clientRoles `json:"golang-app,omitempty"`
	}
	Claims struct {
		ResourceAccess client `json:"resource_access,omitempty"`
		JTI            string `json:"jti,omitempty"`
	}
)

func NewMiddlewares(keycloak pkg.Keycloak, oidc pkg.OpenID) Middlewares {
	return &middlewaresInst{keycloak: keycloak, oidc: oidc}
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

func (mw *middlewaresInst) AuthenticateOIDC(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		if token != "" {
			idToken, err := mw.oidc.Verifier.Verify(ctx.Request().Context(), token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			var IDTokenClaims Claims
			if err := idToken.Claims(&IDTokenClaims); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			fmt.Println(IDTokenClaims)

			return next(ctx)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}
	}
}
