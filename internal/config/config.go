package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type HttpConfig struct {
	Addr          string
	zitadelDomain string
	zitadelPort   int
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
		zitadelDomain: viper.GetString("ZITADEL_DOMAIN"),
		zitadelPort:   viper.GetInt("ZITADEL_PORT"),
	}
}
