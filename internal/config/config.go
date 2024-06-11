package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type HttpConfig struct {
	Addr          string // http server address
	ZitadelDomain string // zitadel domain
	ZitadelPort   string // zitadel port
	CerobsAddr    string // cerbos address
	OtelCollector string // otel collector address
	OtelEnabled   bool
	Debug         bool
}

func NewHttpConfig() *HttpConfig {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	return &HttpConfig{
		Addr:          viper.GetString("HTTP_ADDR"),
		ZitadelDomain: viper.GetString("ZITADEL_DOMAIN"),
		ZitadelPort:   viper.GetString("ZITADEL_PORT"),
		CerobsAddr:    viper.GetString("CERBOS_ADDR"),
		OtelCollector: viper.GetString("OTEL_COLLECTOR"),
		OtelEnabled:   viper.GetBool("OTEL_ENABLED"),
		Debug:         viper.GetBool("DEBUG"),
	}
}
