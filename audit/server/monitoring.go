package main

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goTemp/globalMonitoring"
	"net/http"
)

const (
	// serviceId numeric service identifier
	serviceId = "5"
	// serviceMetricsPrefix prefix for all metrics related to this service
	serviceMetricsPrefix = "goTemp_"
	// metricsEndPoint is the address where we can scrape metrics
	metricsEndPoint = "/metrics"
	//httpPort is the port where the http server for metrics is listening
	httpPort = ":2112"
)

// runHttp runs a secondary server to handle metrics scraping
func runHttp() {
	http.Handle(metricsEndPoint, promhttp.Handler())
	http.ListenAndServe(httpPort, nil)
}

//newMetricsWrapper Create a new metrics wrapper to configure the data to be scraped for monitoring
func newMetricsWrapper() server.HandlerWrapper {
	// TODO: get version number from external source
	return globalMonitoring.NewMetricsWrapper(
		globalMonitoring.ServiceName(serviceName),
		globalMonitoring.ServiceID(serviceId),
		globalMonitoring.ServiceVersion("latest"),
		globalMonitoring.ServiceMetricsPrefix(serviceMetricsPrefix),
		globalMonitoring.ServiceMetricsLabelPrefix(serviceMetricsPrefix),
	)
}

// newMetricsSubscriberWrapper Create a new metrics wrapper to configure the data to be scraped for monitoring when
// receiving a message from the broker
func newMetricsSubscriberWrapper() func(fn broker.Handler) broker.Handler {
	// TODO: get version number from external source
	return globalMonitoring.NewMetricsSubscriberWrapper(
		globalMonitoring.ServiceName(serviceName),
		globalMonitoring.ServiceID(serviceId),
		globalMonitoring.ServiceVersion("latest"),
		globalMonitoring.ServiceMetricsPrefix(serviceMetricsPrefix),
		globalMonitoring.ServiceMetricsLabelPrefix(serviceMetricsPrefix),
	)
}
