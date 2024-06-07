package main

import (
	"context"
	"flow-editor-server/app/account"
	"flow-editor-server/app/flow"
	"flow-editor-server/internal/config"
	"flow-editor-server/internal/validate"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
	"go.uber.org/fx"
	ghttp "goa.design/goa/v3/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fx.New(
		account.Module,
		flow.Module,

		fx.Provide(
			ghttp.NewMuxer,
			validate.NewTranslator,
			validate.NewValidator,
			config.NewHttpConfig,
			newHttpServer,
			newGormDB,
			newZitadel,
			// newCasbinEnforcer,
			newCerbosClient,
		),
		fx.Invoke(func(s *http.Server) {}),
	).Run()
}

// func newCasbinEnforcer(config *config.HttpConfig) (*casbin.Enforcer, error) {
// 	e, err := casbin.NewEnforcer("policy/casbin/model.conf", "policy/casbin/policy.csv")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to casbin: %w", err)
// 	}
// 	return e, nil
// }

func newZitadel(config *config.HttpConfig) (*authorization.Authorizer[*oauth.IntrospectionContext], error) {
	z, err := authorization.New(
		context.Background(),
		zitadel.New(config.ZitadelDomain, zitadel.WithInsecure(config.ZitadelPort)),
		oauth.DefaultAuthorization("key.json"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to zitadel: %w", err)
	}
	return z, nil
}

func newGormDB(lc fx.Lifecycle) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.AutoMigrate(
				&flow.Flow{},
				&account.Account{},
			); err != nil {
				return fmt.Errorf("failed to migrate database: %w", err)
			}
			return nil
		},
	})
	return db, nil
}

func newHttpServer(mux ghttp.ResolverMuxer, config *config.HttpConfig, lc fx.Lifecycle) *http.Server {
	srv := &http.Server{
		Addr:    config.Addr,
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", config.Addr)
			if err != nil {
				return err
			}
			log.Info().Str("addr", config.Addr).Msg("HTTP server started")
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func newCerbosClient() (*cerbos.GRPCClient, error) {
	c, err := cerbos.New("localhost:3593", cerbos.WithPlaintext())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cerbos: %w", err)
	}
	return c, err
}
