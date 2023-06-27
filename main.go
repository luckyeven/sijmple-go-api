package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"fmt"
	"net/http"
)

var (
	example_counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "example_counter",
		Help: "The total number of generated messages",
	})
)

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(example_counter)

	http.HandleFunc("/", HelloServer)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	example_counter.Inc()
	fmt.Fprintf(w, "The simplest API ever!")
}
