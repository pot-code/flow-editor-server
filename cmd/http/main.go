package main

import (
	"openapi-go-demo/app/flow"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
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
	server.HTTPErrorHandler = func(err error, c echo.Context) {
		switch te := err.(type) {
		case *echo.HTTPError:
			c.JSON(te.Code, map[string]string{
				"message": te.Error(),
			})
		case validation.Errors:
			c.JSON(400, te)
		default:
			log.Error().Err(err).Msg("http internal error")
			c.JSON(500, map[string]string{
				"message": err.Error(),
			})
		}
	}

	flow.RegisterHandlers(server, flow.NewController(flow.NewService(db)))

	log.Fatal().Err(server.Start(":3000")).Msg("")
}
