package rest

import (
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/metrics"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/metrics/prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Add metric implementation to context
		ctx := metrics.WithMetrics(r.Context(), prometheus.NewMetrics())

		// Writer wrapper
		ww := middleware.NewWrapResponseWriter(w, 0)

		// Call the next handler in the chain
		next.ServeHTTP(ww, r.WithContext(ctx))

		// Random sleep time to register some more realistic time. Due to the in memory
		// implementation of storage/lock and the lack of external services the response
		// is less than 1ms. Add between 10-60ms delay
		delay := rand.Intn(50) + 10
		time.Sleep(time.Millisecond * time.Duration(delay))

		// Get router, method, status code & response time
		chiCtx := chi.RouteContext(r.Context())
		method := r.Method
		route := getRoutePattern(chiCtx)
		code := ww.Status()
		duration := int(time.Now().Sub(start).Milliseconds())

		// Request metric
		metrics.Request(ctx, route, method, code, duration)
	})
}

func getRoutePattern(chiCtx *chi.Context) string {
	// Extract route pattern from chi context
	route := strings.Replace(strings.Join(chiCtx.RoutePatterns, ""), "/*", "", -1)
	route = strings.Replace(route, "{", "", -1)
	route = strings.Replace(route, "}", "", -1)
	return route
}
