package _logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

// ZapGormLogger @see：https://github.com/go-gorm/gorm/blob/master/logger/logger.go
type ZapGormLogger struct {
	logger                              *zap.Logger
	slowThreshold                       time.Duration // 慢查询阈值
	msgStr                              string
	traceStr, traceWarnStr, traceErrStr string
}

func NewZapGormLogger(zapLogger *zap.Logger, slowThreshold time.Duration) *ZapGormLogger {
	var (
		msgStr       = "%s\n "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)
	return &ZapGormLogger{
		logger:        zapLogger,
		slowThreshold: slowThreshold,
		msgStr:        msgStr,
		traceStr:      traceStr,
		traceWarnStr:  traceWarnStr,
		traceErrStr:   traceErrStr,
	}
}

func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Info(fmt.Sprintf(l.msgStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Warn(fmt.Sprintf(l.msgStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Error(fmt.Sprintf(l.msgStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)

	switch {
	case err != nil && l.logger.Level() <= zap.ErrorLevel:
		sql, rows := fc()
		if rows == -1 {
			l.logger.Error(fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.logger.Error(fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logger.Level() <= zap.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
		if rows == -1 {
			l.logger.Warn(fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.logger.Warn(fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	case l.logger.Level() == zap.DebugLevel:
		sql, rows := fc()
		if rows == -1 {
			l.logger.Debug(fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.logger.Debug(fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	}
}
