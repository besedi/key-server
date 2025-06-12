package srv

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/besedi/key-server/internal/metrics"
)

func withRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func Serve(size int, port int) {
	metrics.Init(size)

	// handlers
	http.Handle("GET /key/{len}", metrics.WithMetrics(withRecovery(KeyHandler(size))))
	http.Handle("GET /key/", metrics.WithMetrics(withRecovery(DefaultHandler(size))))
	http.Handle("/metrics", metrics.MetricsHandler())
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := "0.0.0.0:" + strconv.Itoa(port)
	fmt.Println("Max key size is " + strconv.Itoa(size))
	fmt.Println("Starting server on: " + srv)
	log.Fatal(http.ListenAndServe(srv, nil))
}

func DefaultHandler(size int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/key/"+strconv.Itoa(size), http.StatusFound)
	}
}

func KeyHandler(size int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		len := r.PathValue("len")
		l, err := strconv.Atoi(len)
		if err != nil || l <= 0 || l > size {
			http.Error(w, "Invalid length. Length should not greater then "+strconv.Itoa(size), http.StatusBadRequest)
			return
		}
		metrics.KeyLengthHistogram.Observe(float64(l))
		b := make([]byte, l)
		rand.Read(b)
		for _, bin := range b {
			fmt.Fprintf(w, "%08b", bin)
		}
	}
}
