package config

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
)

type HttpConfig struct {
	Addr            string // http server address
	OidcProvider    string // zitadel domain
	OidcJwkProvider string // zitadel domain
	CerobsAddr      string // cerbos address
	Debug           bool
}

func NewHttpConfig() (*HttpConfig, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	oidc := viper.GetString("OIDC_PROVIDER")
	jwk, err := url.JoinPath(oidc, "/jwks")
	if err != nil {
		return nil, fmt.Errorf("failed to join jwk path: %w", err)
	}
	return &HttpConfig{
		Addr:            viper.GetString("HTTP_ADDR"),
		OidcProvider:    oidc,
		OidcJwkProvider: jwk,
		CerobsAddr:      viper.GetString("CERBOS_ADDR"),
		Debug:           viper.GetBool("DEBUG"),
	}, nil
}
