package main

import (
	"flow-editor-server/app/account"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	if err := db.Where("name=?", "admin").Attrs(&account.Role{Name: "admin"}).FirstOrCreate(&account.Role{}).Error; err != nil {
		log.Fatal().Err(err).Msg("failed to seed database")
	}
	if err := db.Where("name=?", "user").Attrs(&account.Role{Name: "user"}).FirstOrCreate(&account.Role{}).Error; err != nil {
		log.Fatal().Err(err).Msg("failed to seed database")
	}
}
