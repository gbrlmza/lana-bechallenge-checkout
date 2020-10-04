package metrics

import (
	"context"
)

const (
	ctxKey = "metrics"
)

type Tag map[string]interface{}

type Metrics interface {
	Counter(name string, value float64)
	Request(path string, method string, statusCode int, duration int)
}

func WithMetrics(ctx context.Context, metrics Metrics) context.Context {
	return context.WithValue(ctx, ctxKey, metrics)
}

func Counter(ctx context.Context, name string, value float64) {
	if metric, ok := ctx.Value(ctxKey).(Metrics); ok {
		metric.Counter(name, value)
	}
}

func Request(ctx context.Context, path string, method string, statusCode int, duration int) {
	if metric, ok := ctx.Value(ctxKey).(Metrics); ok {
		metric.Request(path, method, statusCode, duration)
	}
}
