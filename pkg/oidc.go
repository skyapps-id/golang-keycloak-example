package pkg

import (
	"context"

	oidc "github.com/coreos/go-oidc"
)

type (
	OpenID struct {
		Provider   *oidc.Provider // keycloak client
		OidcConfig *oidc.Config   // ClientID specified in Keycloak
		Verifier   *oidc.IDTokenVerifier
		// ClientSecret string // client secret specified in Keycloak
		// Realm        string // realm specified in Keycloak
	}
)

func NewOpenID() OpenID {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/golang-svc")
	if err != nil {
		panic(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: "golang-app",
	}
	verifier := provider.Verifier(oidcConfig)

	return OpenID{
		Provider:   provider,
		OidcConfig: oidcConfig,
		Verifier:   verifier,
		// ClientSecret: "Y2erLgUO98pqBJGyMW5BdEEXPyawWowj",
	}
}
