package orm

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}
	return db, nil
}
