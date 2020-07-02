package config

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"agokit/i18n"
	"os"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	"github.com/lightstep/lightstep-tracer-go"
	"sourcegraph.com/sourcegraph/appdash"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config interface {
	Load() (error)
}

//New Config is  to set params
func New(basePath string) Config {
	c := BasicConfig{}
	flag.StringVar(&c.arg.ConfigPath, "config-path", fmt.Sprintf("%sconfigs", basePath), "Config Path")
	flag.StringVar(&c.arg.LanguagePath, "language-path", fmt.Sprintf("%slanguages", basePath), "Language Path")
 	c.app.BasePath = basePath
	return &c
}

type BasicConfig struct {
	arg Arg
	app App
	com Com
	mws Middlewares
}

//Load config from app.yml file
func (c *BasicConfig) Load() ( error ){
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.arg.ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&c.app)

	file := fmt.Sprintf("%s-com.yml",c.app.Env )
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.arg.ConfigPath)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&c.com)
	return  err
}

//Load config  for middlewares
func (c *BasicConfig) LoadMiddlewares() (mws Middlewares) {
	fieldKeys := []string{"method", "error"}
	mws.RequestCount = kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: c.app.Name,
		Subsystem: c.app.Service,
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	mws.RequestLatency = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: c.app.Name,
		Subsystem: c.app.Service,
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	mws.CountResult = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: c.app.Name,
		Subsystem: c.app.Service,
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: c.app.Name,
			Subsystem: c.app.Service,
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}
	mws.Duration = duration

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	mws.Logger = logger
	cfg := c
	zipkinURL := fmt.Sprintf("http://%s:%d%s", cfg.com.Zipkin.Host, cfg.com.Zipkin.Port, cfg.com.Zipkin.Api)
	mws.ZipkinURL = zipkinURL

	var zipkinTracer *zipkin.Tracer
	{
		if zipkinURL != "" {
			var (
				err      error

				hostPort = fmt.Sprintf("%s:%d", cfg.com.Zipkin.Host, cfg.com.Zipkin.Port)
				reporter = zipkinhttp.NewReporter(zipkinURL)
			)

			logger.Log("Servicename", cfg.app.Service)

			mws.Reporter = reporter
			zEP, _ := zipkin.NewEndpoint(cfg.app.Service, hostPort)
			zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP))
			if err != nil {
				logger.Log("err", err)
				os.Exit(1)
			}
			if !(cfg.com.Zipkin.Bridge) {
				logger.Log("tracer", "Zipkin", "type", "Native", "URL", zipkinURL)
			}
		}
	}
	mws.ZipkinTracer = zipkinTracer

	// Determine which OpenTracing tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency.
	var tracer stdopentracing.Tracer
	{
		if cfg.com.Zipkin.Bridge && zipkinTracer != nil {
			logger.Log("tracer", "Zipkin", "type", "OpenTracing", "URL", zipkinURL)
			tracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil // do not instrument with both native tracer and opentracing bridge
		} else if cfg.com.Lightstep.Token != "" {
			logger.Log("tracer", "LightStep") // probably don't want to print out the token :)
			lightStepTracer := lightstep.NewTracer(lightstep.Options{
				Collector:   lightstep.Endpoint{},
				AccessToken: cfg.com.Lightstep.Token,
				Tags: map[string]interface{}{
					lightstep.ComponentNameKey: cfg.app.Service,
					"service.version":          cfg.app.Version,
				},
			})
			stdopentracing.SetGlobalTracer(lightStepTracer)
			tracer = stdopentracing.GlobalTracer()
			//span := tracer.StartSpan("my-first-span")
			//span.SetTag("kind", "server")
			//span.LogKV("message", "what a lovely day")
			//span.Finish()

			mws.LightstepTracer = lightStepTracer

		} else if cfg.com.Appdashot.Port != 0 {
			logger.Log("tracer", "Appdash", "addr", cfg.com.Appdashot.Port)
			apphost := fmt.Sprintf("%s:%d", cfg.com.Appdashot.Host, cfg.com.Appdashot.Port)
			tracer = appdashot.NewTracer(appdash.NewRemoteCollector(apphost))
		} else {
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}

	mws.Tracer = tracer

	mws.Locale = i18n.New(c.arg.LanguagePath, c.app.AcceptLanguage)

	c.mws = mws

	return mws
}

func (c *BasicConfig) LoadDB() (db *gorm.DB, err error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%s&charset=%s",
		c.com.DB.Username, c.com.DB.Password,
		c.com.DB.Host, c.com.DB.Port,
		c.com.DB.Database, c.com.DB.ParseTime, c.com.DB.Charset)
	db, err = gorm.Open("mysql", dataSource)
	return db, err
}