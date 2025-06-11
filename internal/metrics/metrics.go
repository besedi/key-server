package metrics

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics definitions
var (
	statusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_status_codes_total",
			Help: "Count of HTTP status codes returned",
		},
		[]string{"code"},
	)

	KeyLengthHistogram prometheus.Histogram
)

// Setup registers metrics and initializes histogram buckets
func Init(maxSize int) {
	step := float64(maxSize) / 20
	buckets := make([]float64, 20)
	for i := 0; i < 20; i++ {
		buckets[i] = step * float64(i+1)
	}
	KeyLengthHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "key_length_distribution",
		Help:    "Histogram of requested key lengths",
		Buckets: buckets,
	})

	prometheus.MustRegister(statusCounter)
	prometheus.MustRegister(KeyLengthHistogram)
}

// Middleware to track HTTP status codes
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func WithMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		h.ServeHTTP(rec, r)
		statusCounter.WithLabelValues(strconv.Itoa(rec.status)).Inc()
	})
}

// Expose metrics endpoint
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
