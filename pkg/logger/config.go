package logger

import (
	"path/filepath"

	"backend/config"
)

type loggerConfig struct {
	LogFile      string
	MaxSize      int
	MaxBackups   int
	MaxAge       int
	Compress     bool
	EnableFile   bool
	EnableStdout bool
	TimeZone     string
}

func NewConfig(srvCfg *config.Server) *loggerConfig {
	return &loggerConfig{
		LogFile:      filepath.Join("logs", "app.log"),
		MaxSize:      10,
		MaxBackups:   3,
		MaxAge:       7,
		Compress:     false,
		EnableStdout: srvCfg.EnableStdout,
		EnableFile:   srvCfg.EnableFile,
		TimeZone:     srvCfg.TimeZone,
	}
}

// TODO: read more config from file, env,...
