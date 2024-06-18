package log

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct{}

func NewGormLogger() *GormLogger {
	return &GormLogger{}
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *GormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	InfoF(ctx, msg, args...)
}

func (g *GormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	WarnF(ctx, msg, args...)
}

func (g *GormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	ErrorF(ctx, msg, args...)
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	cost := time.Now().Sub(begin).Milliseconds()
	if err == nil {
		InfoF(ctx, "mysql success, sql:[%v] rows:[%v] cost:[%v]", sql, rows, cost)
	} else {
		ErrorF(ctx, "mysql fail, sql:[%v] rows:[%v] cost:[%v] err:[%]", sql, rows, cost, err)
	}
}
