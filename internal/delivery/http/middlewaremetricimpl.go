package http

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricMiddlewareImpl struct {
	httpRequestsTotal   *prometheus.CounterVec
	httpRequestDuration *prometheus.HistogramVec
}

var _ MetricMiddleware = &MetricMiddlewareImpl{}

func ProvideMetricMiddleware() *MetricMiddlewareImpl {
	return &MetricMiddlewareImpl{
		httpRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		httpRequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "path", "status"},
		),
	}
}

func (m *MetricMiddlewareImpl) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(ctx.Writer.Status())
		path := ctx.FullPath()
		if path == "" {
			path = ctx.Request.URL.Path
		}

		m.httpRequestsTotal.WithLabelValues(ctx.Request.Method, path, status).Inc()
		m.httpRequestDuration.WithLabelValues(ctx.Request.Method, path, status).Observe(duration)
	}
}
