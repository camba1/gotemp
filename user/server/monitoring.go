package main

import (
	"github.com/micro/go-micro/v2/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goTemp/globalMonitoring"
	"net/http"
)

const (
	// serviceId numeric service identifier
	serviceId = "1"
	// serviceMetricsPrefix prefix for all metrics related to this service
	serviceMetricsPrefix = "goTemp_"
)

// runHttp runs a secondary server to handle metrics scraping
func runHttp() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

//newMetricsWrapper Create a new metrics wrapper to configreu the data to be scraped for monitoring
func newMetricsWrapper() server.HandlerWrapper {
	return globalMonitoring.NewMetricsWrapper(
		globalMonitoring.ServiceName(serviceName),
		globalMonitoring.ServiceID(serviceId),
		globalMonitoring.ServiceVersion("latest"),
		globalMonitoring.ServiceMetricsPrefix(serviceMetricsPrefix),
		globalMonitoring.ServiceMetricsLabelPrefix(serviceMetricsPrefix),
	)
}
