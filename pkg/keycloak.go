package pkg

import "github.com/Nerzal/gocloak/v11"

type (
	Keycloak struct {
		Gocloak      gocloak.GoCloak // keycloak client
		ClientID     string          // ClientID specified in Keycloak
		ClientSecret string          // client secret specified in Keycloak
		Realm        string          // realm specified in Keycloak
	}
)

func NewKeycloak() Keycloak {
	return Keycloak{
		Gocloak:      gocloak.NewClient("http://localhost:8080"),
		ClientID:     "golang-app",
		ClientSecret: "Y2erLgUO98pqBJGyMW5BdEEXPyawWowj",
		Realm:        "golang-svc",
	}
}
