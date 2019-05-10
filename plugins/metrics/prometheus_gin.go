/*
@Time : 2019/5/9 17:56 
@Author : kenny zhu
@File : prometheus.go
@Software: GoLand
@Others: gin prometheus process
*/
package metrics

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

var (
	defaultGinMetricPrefix = "gin"
	DefaultPrometheus = NewGinHandlerWrapper()
)

type Option func(o *Options)

// for gin wrapper, support goroutine
func NewGinHandlerWrapper(opts ...Option) gin.HandlerFunc {
	md := make(map[string]string)
	gopts := Options{}

	for _, opt := range opts {
		opt(&gopts)
	}

	for k, v := range gopts.MetaData {
		md[fmt.Sprintf("%s_%s", defaultGinMetricPrefix, k)] = v
	}
	if len(gopts.Namespace) > 0 {
		md[fmt.Sprintf("%s_%s", defaultGinMetricPrefix, "name")] = gopts.Namespace
	}
	if len(gopts.Id) > 0 {
		md[fmt.Sprintf("%s_%s", defaultGinMetricPrefix, "id")] = gopts.Id
	}
	if len(gopts.Version) > 0 {
		md[fmt.Sprintf("%s_%s", defaultGinMetricPrefix, "version")] = gopts.Version
	}

	// counter calls directly
	opsCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: gopts.Namespace,
			Name:      "request_total",
			Help:      "How many gin service requests processed, partitioned by method and status",
		},
		[]string{"method", "status"},
	)

	timeCounterSummary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: gopts.Namespace,
			Name:      "upstream_latency_microseconds",
			Help:      "Service backend method request latencies in microseconds",
		},
		[]string{"method"},
	)

	timeCounterHistogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: gopts.Namespace,
			Name:      "request_duration_seconds",
			Help:      "Service method request time in seconds",
		},
		[]string{"method"},
	)

	reg := prometheus.NewRegistry()
	wrapreg := prometheus.WrapRegistererWith(md, reg)
	wrapreg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{ Namespace:gopts.Namespace }),
		prometheus.NewGoCollector(),
		opsCounter,
		timeCounterSummary,
		timeCounterHistogram,
	)

	prometheus.DefaultGatherer = reg
	prometheus.DefaultRegisterer = wrapreg

	return func(ctx *gin.Context) {
		// name here may be the domain name if dns query used
		var name string
		if ctx.Request.Host != "" {
			name = ctx.Request.Host
		} else {
			name = ctx.Request.URL.Host
		}

		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			us := v * 1000000 // make microseconds
			timeCounterSummary.WithLabelValues(name).Observe(us)
			timeCounterHistogram.WithLabelValues(name).Observe(v)
		}))
		defer timer.ObserveDuration()

		// call and judge the result.
		ctx.Next()
		if statusCode := ctx.Writer.Status(); statusCode >= http.StatusMultipleChoices {
			opsCounter.WithLabelValues(name, "success").Inc()
		} else {
			opsCounter.WithLabelValues(name, "fail").Inc()
		}
	}
}