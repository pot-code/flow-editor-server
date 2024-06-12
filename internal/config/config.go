package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type HttpConfig struct {
	Addr            string // http server address
	OidcProvider    string // zitadel domain
	OidcJwkProvider string // zitadel domain
	OidcApiID       string // zitadel domain
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

	return &HttpConfig{
		Addr:            viper.GetString("HTTP_ADDR"),
		OidcProvider:    viper.GetString("OIDC_PROVIDER"),
		OidcJwkProvider: viper.GetString("OIDC_JWK_PROVIDER"),
		OidcApiID:       viper.GetString("OIDC_API_ID"),
		CerobsAddr:      viper.GetString("CERBOS_ADDR"),
		Debug:           viper.GetBool("DEBUG"),
	}, nil
}
