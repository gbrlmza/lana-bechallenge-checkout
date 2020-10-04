package prometheus

import (
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"sync"
)

var (
	counters    map[string]prometheus.Counter
	once        sync.Once
	httpRequest *prometheus.HistogramVec
)

// Implementation of Metrics interface with Prometheus
type Prometheus struct{}

func NewMetrics() metrics.Metrics {
	once.Do(func() {
		// Map counter dynamically created counter
		counters = make(map[string]prometheus.Counter)

		// Standard request
		httpRequest = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Subsystem: "http",
			Name:      "request_duration_milliseconds",
			Help:      "The latency of the HTTP requests.",
		}, []string{"Handler", "Method", "StatusCode"})
		prometheus.MustRegister(httpRequest)
	})

	return &Prometheus{}
}

func (p *Prometheus) Counter(name string, value float64) {
	// Check if counter exists. Create it if new
	if _, exists := counters[name]; !exists {
		counters[name] = promauto.NewCounter(prometheus.CounterOpts{Name: name})
	}

	// Increment
	counters[name].Add(value)
}

func (p *Prometheus) Request(path string, method string, statusCode int, duration int) {
	httpRequest.With(prometheus.Labels{
		"Handler":    path,
		"Method":     method,
		"StatusCode": strconv.Itoa(statusCode),
	}).Observe(float64(duration))
}
