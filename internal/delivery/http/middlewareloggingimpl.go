package http

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggingMiddlewareImpl struct {
	logger *zap.Logger
}

var _ LoggingMiddleware = &LoggingMiddlewareImpl{}

func ProvideLoggingMiddleware(logger *zap.Logger) *LoggingMiddlewareImpl {
	return &LoggingMiddlewareImpl{
		logger: logger,
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (m *LoggingMiddlewareImpl) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		status := ctx.Writer.Status()
		fields := []zapcore.Field{
			zap.Int("status", status),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.Duration("latency", latency),
		}
		if status >= 400 {
			fields = append(fields, zap.String("response_body", blw.body.String()))
		}

		if len(ctx.Errors) > 0 && status >= 500 {
			for _, e := range ctx.Errors.Errors() {
				m.logger.Error(e, fields...)
			}
			return
		}
		m.logger.Info(path, fields...)
	}
}
