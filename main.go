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
		Help: "An example of a Prometheus counter",
	})
)

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(example_counter)

	http.HandleFunc("/", Handler)
	http.Handle("/metrics", promhttp.Handler()) // Exporting metrics
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	example_counter.Inc()
	fmt.Println("Handler executing")
	fmt.Fprintf(w, "The simplest API ever!")
}
