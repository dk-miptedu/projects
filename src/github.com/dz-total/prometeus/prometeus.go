package prometeus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests.",
		},
		[]string{"path"}, // labels
	)
)

func Init() {
	// Регистрация метрики в Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometeusHanler()
}

func prometeusHanler() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8089", nil)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	// Увеличение счетчика для метрик
	httpRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	w.Write([]byte("Hello, world!"))
}
