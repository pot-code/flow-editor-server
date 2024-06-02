package main

import (
	"context"
	"errors"
	"openapi-go-demo/app/account"
	"openapi-go-demo/app/flow"
	"openapi-go-demo/middleware"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	zitadel, err := authorization.New(
		context.Background(),
		zitadel.New(viper.GetString("ZITADEL_DOMAIN"), zitadel.WithInsecure(viper.GetString("ZITADEL_PORT"))),
		oauth.DefaultAuthorization("key.json"),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create zitadel")
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	if err := db.AutoMigrate(&flow.FlowModel{}, &account.AccountModel{}); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate database")
	}

	server := echo.New()
	server.HTTPErrorHandler = func(err error, c echo.Context) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, echo.Map{
				"message": err.Error(),
			})
			return
		}

		switch te := err.(type) {
		case *echo.HTTPError:
			c.JSON(te.Code, echo.Map{
				"message": te.Error(),
			})
		case validation.Errors:
			c.JSON(400, te)
		default:
			log.Error().Err(err).Msg("http internal error")
			c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}
	}
	server.Use(middleware.OauthInterceptor(zitadel))

	flow.RegisterHandlers(server, flow.NewController(flow.NewService(db)))
	account.RegisterHandlers(server, account.NewController(db))

	log.Fatal().Err(server.Start(":3000")).Msg("")
}
