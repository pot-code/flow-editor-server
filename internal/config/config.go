package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type HttpConfig struct {
	Addr          string // http server address
	ZitadelDomain string // zitadel domain
	ZitadelPort   string // zitadel port
	CerobsAddr    string // cerbos address
	Debug         bool
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
		Addr:          viper.GetString("HTTP_ADDR"),
		ZitadelDomain: viper.GetString("ZITADEL_DOMAIN"),
		ZitadelPort:   viper.GetString("ZITADEL_PORT"),
		CerobsAddr:    viper.GetString("CERBOS_ADDR"),
		Debug:         viper.GetBool("DEBUG"),
	}, nil
}
