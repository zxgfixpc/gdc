package logger

import (
	"fmt"
	"os"
	"time"

	"_gdc_/conf"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogger *zap.SugaredLogger

const (
	logLevelInfo = "info"
	logLevelErr  = "error"
)

func initZapLog() error {
	logConf := &conf.LogConf{}
	if err := conf.Parser(logConf, conf.LogConfPath); err != nil {
		return err
	}

	// zap core:encoder writer level
	cores := make([]zapcore.Core, 0, 2)
	for _, zConf := range logConf.Zap {
		level := zapcore.DebugLevel
		if zConf.Level == logLevelErr {
			level = zapcore.ErrorLevel
		}
		encoder := getEncoder()
		writeSyncer := getLogWriter(zConf)
		item := zapcore.NewCore(encoder, writeSyncer, level)
		cores = append(cores, item)
	}
	core := zapcore.NewTee(cores...)

	zapLogger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1), // 调用栈-1，log日志包了一层
	).Sugar()

	return nil
}

func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encoderConfig := zap.NewProductionEncoderConfig()
	// 序列化时间。eg: 2024-06-11T19:11:35.921+0800
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logConf conf.ZapLogConf) zapcore.WriteSyncer {
	// zap log不支持日志文件切割，使用Lumberjack
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf(logConf.Filename, time.Now().Unix()), // 日志文件的位置
		MaxSize:    logConf.MaxSize,                                  // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: logConf.MaxBackups,                               // 保留旧文件的最大个数
		MaxAge:     logConf.MaxAge,                                   // 保留旧文件的最大天数
		Compress:   false,                                            // 是否压缩/归档旧文件
	}

	syncFile := zapcore.AddSync(lumberJackLogger)
	syncConsole := zapcore.AddSync(os.Stderr) // 打印到控制台
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}

func Shutdown() {
	_ = zapLogger.Sync()
}
