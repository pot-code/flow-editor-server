package main

import (
	"openapi-go-demo/app/flow"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	if err := db.AutoMigrate(&flow.FlowModel{}); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate database")
	}

	server := echo.New()
	flow.RegisterHandlers(server, flow.NewController(flow.NewService(db)))

	log.Fatal().Err(server.Start(":3000")).Msg("")
}
