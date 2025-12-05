package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(cfg *loggerConfig) *zap.Logger {
	if !cfg.EnableStdout && !cfg.EnableFile {
		return zap.NewNop()
	}

	var cores []zapcore.Core
	logLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		loc, _ := time.LoadLocation(cfg.TimeZone)
		enc.AppendString(t.In(loc).Format(time.RFC3339))
	}
	if cfg.EnableStdout {
		stdout := zapcore.AddSync(os.Stdout)
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = timeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		cores = append(cores, zapcore.NewCore(consoleEncoder, stdout, logLevel))
	}

	if cfg.EnableFile {
		file := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.LogFile,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		})

		cfg := zap.NewProductionEncoderConfig()
		cfg.TimeKey = "timestamp"
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder

		fileEncoder := zapcore.NewJSONEncoder(cfg)
		cores = append(cores, zapcore.NewCore(fileEncoder, file, logLevel))
	}
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return logger
}
