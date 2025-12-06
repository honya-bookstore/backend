package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateRedisGetField(id, key string, err error) *[]zapcore.Field {
	return &[]zapcore.Field{
		zap.String("id", id),
		zap.String("cache_key", key),
		zap.String("error", err.Error()),
	}
}

func CreateRedisListField(key string, err error) *[]zapcore.Field {
	return &[]zapcore.Field{
		zap.String("cache_key", key),
		zap.String("error", err.Error()),
	}
}
