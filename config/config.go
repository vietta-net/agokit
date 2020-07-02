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
	"github.com/vietta-net/agokit"
	"os"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	"github.com/lightstep/lightstep-tracer-go"
	"sourcegraph.com/sourcegraph/appdash"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config interface {
	LoadDB() (db *gorm.DB, err error)
	LoadMiddlewares() (mws Middlewares)
}

//New Config is  to set params
func New(basePath string) (Config , error){
	c := BasicConfig{}
	flag.StringVar(&c.Arg.ConfigPath, "config-path", fmt.Sprintf("%sconfigs", basePath), "Config Path")
	flag.StringVar(&c.Arg.LanguagePath, "language-path", fmt.Sprintf("%slanguages", basePath), "Language Path")
 	c.App.BasePath = basePath
	err:= c.Load()
	return &c, err
}

type BasicConfig struct {
	Arg Args
	App Application
	Com Component
	Mws Middlewares
	Bb  *gorm.DB
}

//Load config from app.yml file
func (c *BasicConfig) Load() ( error ){
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.Arg.ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&c.App)

	file := fmt.Sprintf("%s-com.yml",c.app.Env )
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.Arg.ConfigPath)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&c.Com)
	return  err
}

//Load config  for middlewares
func (c *BasicConfig) LoadMiddlewares() (mws Middlewares) {
	fieldKeys := []string{"method", "error"}
	mws.RequestCount = kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: c.App.Name,
		Subsystem: c.App.Service,
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	mws.RequestLatency = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: c.App.Name,
		Subsystem: c.App.Service,
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	mws.CountResult = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: c.App.Name,
		Subsystem: c.App.Service,
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: c.App.Name,
			Subsystem: c.App.Service,
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

			logger.Log("Servicename", cfg.App.Service)

			mws.Reporter = reporter
			zEP, _ := zipkin.NewEndpoint(cfg.App.Service, hostPort)
			zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP))
			if err != nil {
				logger.Log("err", err)
				os.Exit(1)
			}
			if !(cfg.Com.Zipkin.Bridge) {
				logger.Log("tracer", "Zipkin", "type", "Native", "URL", zipkinURL)
			}
		}
	}
	mws.ZipkinTracer = zipkinTracer

	// Determine which OpenTracing tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency.
	var tracer stdopentracing.Tracer
	{
		if cfg.Com.Zipkin.Bridge && zipkinTracer != nil {
			logger.Log("tracer", "Zipkin", "type", "OpenTracing", "URL", zipkinURL)
			tracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil // do not instrument with both native tracer and opentracing bridge
		} else if cfg.Com.Lightstep.Token != "" {
			logger.Log("tracer", "LightStep") // probably don't want to print out the token :)
			lightStepTracer := lightstep.NewTracer(lightstep.Options{
				Collector:   lightstep.Endpoint{},
				AccessToken: cfg.Com.Lightstep.Token,
				Tags: map[string]interface{}{
					lightstep.ComponentNameKey: cfg.App.Service,
					"service.version":          cfg.App.Version,
				},
			})
			stdopentracing.SetGlobalTracer(lightStepTracer)
			tracer = stdopentracing.GlobalTracer()
			//span := tracer.StartSpan("my-first-span")
			//span.SetTag("kind", "server")
			//span.LogKV("message", "what a lovely day")
			//span.Finish()

			mws.LightstepTracer = lightStepTracer

		} else if cfg.Com.Appdashot.Port != 0 {
			logger.Log("tracer", "Appdash", "addr", cfg.Com.Appdashot.Port)
			apphost := fmt.Sprintf("%s:%d", cfg.com.Appdashot.Host, cfg.com.Appdashot.Port)
			tracer = appdashot.NewTracer(appdash.NewRemoteCollector(apphost))
		} else {
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}

	mws.Tracer = tracer

	mws.Locale = i18n.New(c.Arg.LanguagePath, c.App.AcceptLanguage)

	c.Mws = mws

	return mws
}

func (c *BasicConfig) LoadDB() (db *gorm.DB, err error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%s&charset=%s",
		c.com.DB.Username, c.com.DB.Password,
		c.com.DB.Host, c.com.DB.Port,
		c.com.DB.Database, c.com.DB.ParseTime, c.com.DB.Charset)
	db, err = gorm.Open("mysql", dataSource)
	c.Bb = db
	return db, err
}