package logger

import (
	"context"
	"fmt"
	"gdc/lib/trace"
)

func Debug(ctx context.Context, msg string) {
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Debugw(msg, "trace_id", traceID)
}

func DebugF(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Debugw(msg, "trace_id", traceID)
}

func Info(ctx context.Context, msg string) {
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Infow(msg, "trace_id", traceID)
}

func InfoF(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Infow(msg, "trace_id", traceID)
}

func Warn(ctx context.Context, msg string) {
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Warnw(msg, "trace_id", traceID)
}

func WarnF(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Warnw(msg, "trace_id", traceID)
}

func Error(ctx context.Context, msg string) {
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Errorw(msg, "trace_id", traceID)
}

func ErrorF(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Errorw(msg, "trace_id", traceID)
}

func Panic(ctx context.Context, msg string) {
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Panicw(msg, "trace_id", traceID)
}

func PanicF(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	traceID := trace.GetTraceIDByCtx(ctx)
	zapLogger.Panicw(msg, "trace_id", traceID)
}

func InitLog() error {
	return initZapLog()
}
