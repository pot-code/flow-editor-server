package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type HttpConfig struct {
	Addr          string
	ZitadelDomain string
	ZitadelPort   string
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
	}
}
