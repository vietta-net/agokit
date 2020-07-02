package config

import (
	"github.com/go-kit/kit/log"
	"github.com/lightstep/lightstep-tracer-go"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
 	"github.com/go-kit/kit/metrics"
	"agokit/i18n"
)

type Middlewares struct {
	Logger          log.Logger
	ZipkinTracer    *zipkin.Tracer
	ZipkinURL       string
	Reporter        reporter.Reporter
	Tracer          stdopentracing.Tracer
	LightstepTracer lightstep.Tracer
	RequestCount    metrics.Counter
	RequestLatency  metrics.Histogram
	CountResult     metrics.Histogram
	Duration        metrics.Histogram
	Locale			i18n.Locale
}

