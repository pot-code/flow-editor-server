package orm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: &gormLogger{
			l:             log.Logger,
			slowThreshold: 2 * time.Second,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}
	return db, nil
}

type gormLogger struct {
	l             zerolog.Logger
	slowThreshold time.Duration
}

// Error implements logger.Interface.
func (g *gormLogger) Error(ctx context.Context, pattern string, args ...interface{}) {
	g.l.Error().Ctx(ctx).Msgf(pattern, args...)
}

// Info implements logger.Interface.
func (g *gormLogger) Info(ctx context.Context, pattern string, args ...interface{}) {
	g.l.Info().Ctx(ctx).Msgf(pattern, args...)
}

// LogMode implements logger.Interface.
func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	zl := zerolog.Disabled
	switch level {
	case logger.Info:
		zl = zerolog.InfoLevel
	case logger.Warn:
		zl = zerolog.WarnLevel
	case logger.Error:
		zl = zerolog.ErrorLevel
	}
	return &gormLogger{
		l: g.l.Level(zl),
	}
}

// Trace implements logger.Interface.
func (g *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows := fc()
		g.l.Error().Err(err).Int64("rows", rows).Dur("cost", elapsed).Str("sql", sql).Send()
	case elapsed > g.slowThreshold && g.slowThreshold != 0:
		sql, rows := fc()
		g.l.Warn().Int64("rows", rows).Dur("cost", elapsed).Str("sql", sql).Msgf("SLOW SQL >= %v", g.slowThreshold)
	default:
		sql, rows := fc()
		g.l.Debug().Int64("rows", rows).Dur("cost", elapsed).Str("sql", sql).Send()
	}
}

// Warn implements logger.Interface.
func (g *gormLogger) Warn(ctx context.Context, pattern string, args ...interface{}) {
	g.l.Warn().Ctx(ctx).Msgf(pattern, args...)
}

var _ logger.Interface = (*gormLogger)(nil)
