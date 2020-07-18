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
	"github.com/vietta-net/agokit/i18n"
	"net/url"
	"os"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	"github.com/lightstep/lightstep-tracer-go"
	"sourcegraph.com/sourcegraph/appdash"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

type Config interface {
	LoadDB() (db *gorm.DB, err error)
	LoadMiddlewares() (mws Middlewares)
}

//New Config is  to set params
func New(basePath string) (Config , error){
	c := BasicConfig{}



	if os.Getenv("APP_TEST") == ""{
		flag.UintVar(&c.Arg.Mode, "mode", 0, "Mode")
		flag.StringVar(&c.Arg.BasePath, "base-path", basePath, "Base Path")
		flag.StringVar(&c.Arg.ConfigPath, "config-path", fmt.Sprintf("%s/configs", basePath), "Config Path")
		flag.StringVar(&c.Arg.LanguagePath, "language-path", fmt.Sprintf("%s/languages", basePath), "Language Path")
		flag.Parse()
	}else{
		c.Arg.BasePath = basePath
		c.Arg.ConfigPath  = fmt.Sprintf("%s/configs", basePath)
		c.Arg.LanguagePath = fmt.Sprintf("%s/languages", basePath)
	}

	c.App.BasePath = c.Arg.BasePath
	err:= c.Load()

	if os.Getenv("APP_PREFIX") != "" {
		viper.SetEnvPrefix(os.Getenv("APP_PREFIX"))
	}
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	//NAME string
	if str :=  viper.GetString("NAME");  str != "" {
		c.App.Name = str
	}
	//SERVICE  string
	if str :=  viper.GetString("SERVICE");  str != "" {
		c.App.Service = str
	}
	//VERSION string
	if str :=  viper.GetString("VERSION");  str != "" {
		c.App.Version = str
	}
	//MODE  uint8
	if iVar, err := strconv.ParseInt(viper.GetString("MODE"), 10, 16) ;  err == nil && iVar > 0 {
		c.App.Mode = uint8(iVar)
	}
	//TIMEZONE string
	if str :=  viper.GetString("TIMEZONE");  str != "" {
		c.App.Timezone = str
	}
	//ENV string
	if str :=  viper.GetString("ENV");  str != "" {
		c.App.Env = str
	}
	//JWT_SECRET string
	if str :=  viper.GetString("JWT_SECRET");  str != "" {
		c.App.Secret = str
	}
	//BASE_PATH string
	if str :=  viper.GetString("BASE_PATH");  str != "" {
		c.App.BasePath = str
		c.Arg.BasePath = str
	}
	//CONFIG_PATH string
	if str :=  viper.GetString("CONFIG_PATH");  str != "" {
		c.Arg.ConfigPath = str
	}
	//LANGUAGE_PATH string
	if str :=  viper.GetString("LANGUAGE_PATH");  str != "" {
		c.Arg.LanguagePath = str
	}
	//DB_DIALECT  string
	if str :=  viper.GetString("DB_DIALECT");  str != "" {
		c.Com.DB.Dialect = str
	}
	//DB_HOST string
	if str :=  viper.GetString("DB_HOST");  str != "" {
		c.Com.DB.Host = str
	}
	//DB_PORT uint16
	if iVar, err := strconv.ParseUint(viper.GetString("DB_PORT"), 10, 64) ;  err == nil && iVar > 0 {
		c.Com.DB.Port = uint16(iVar)
	}
	//DB_DATABASE string
	if str :=  viper.GetString("DB_DATABASE");  str != "" {
		c.Com.DB.Database = str
	}
	//DB_USER string
	if str :=  viper.GetString("DB_USER");  str != "" {
		c.Com.DB.Username = str
	}
	//DB_PASSWORD string
	if str :=  viper.GetString("DB_PASSWORD");  str != "" {
		c.Com.DB.Password = str
	}
	//GRPC_PORT uint16
	if iVar, err := strconv.ParseUint(viper.GetString("GRPC_PORT"), 10, 64) ;  err == nil && iVar > 0 {
		c.Com.Grpc.Port = uint16(iVar)
	}

	//DEBUG_PORT uint16
	if iVar, err := strconv.ParseUint(viper.GetString("DEBUG_PORT"), 10, 64) ;  err == nil && iVar > 0 {
		c.Com.Debug.Port = uint16(iVar)
	}
	//REST_PORT  uint16
	if iVar, err := strconv.ParseUint(viper.GetString("REST_PORT"), 10, 64) ;  err == nil && iVar > 0 {
		c.Com.Rest.Port = uint16(iVar)
	}
	//REST_BASE_URL string
	if str :=  viper.GetString("REST_BASE_URL");  str != "" {
		c.Com.Rest.BaseUrl = str
	}
	//ZIPKIN_HOST string
	if str :=  viper.GetString("ZIPKIN_HOST");  str != "" {
		c.Com.Zipkin.Host = str
	}
	//ZIPKIN_API string
	if str :=  viper.GetString("ZIPKIN_API");  str != "" {
		c.Com.Zipkin.Api = str
	}
	//ZIPKIN_PORT uint16
	if iVar, err := strconv.ParseUint(viper.GetString("ZIPKIN_PORT"), 10, 64) ;  err == nil && iVar > 0 {
		c.Com.Zipkin.Port = uint16(iVar)
	}
	//ZIPKIN_BRIDGE bool
	if bVar, err := strconv.ParseBool(viper.GetString("ZIPKIN_BRIDGE")) ;  err == nil && bVar {
		c.Com.Zipkin.Bridge = bVar
	}
	//LIGHTSTEP_TOKEN string
	if str :=  viper.GetString("LIGHTSTEP_TOKEN");  str != "" {
		c.Com.Lightstep.Token = str
	}
	//APPDASHOT_HOST string
	if str :=  viper.GetString("APPDASHOT_HOST");  str != "" {
		c.Com.Appdashot.Host = str
	}
	//APPDASHOT_PORT uint16
	if iVar, err := strconv.ParseUint(viper.GetString("APPDASHOT_PORT"), 10, 64) ;  err == nil && iVar > 0 {
		c.Com.Appdashot.Port = uint16(iVar)
	}

	return &c, err
}

type BasicConfig struct {
	Arg Args
	App Application
	Com Component
	Mws Middlewares
	Bb  *gorm.DB
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//Load config from app.yml file
func (c *BasicConfig) Load() ( error ){
	AppFile := fmt.Sprintf("%s/app.yml", c.Arg.ConfigPath)
	if ! FileExists(AppFile) {
		panic(AppFile + " is not existed")
	}
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.Arg.ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c.App)

	file := fmt.Sprintf("%s-com.yml",c.App.Env )
	ComFile := fmt.Sprintf("%s/%s", c.Arg.ConfigPath, file)

	if ! FileExists(ComFile) {
		panic(ComFile + " is not existed")
	}

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
	zipkinURL := fmt.Sprintf("http://%s:%d%s", cfg.Com.Zipkin.Host, cfg.Com.Zipkin.Port, cfg.Com.Zipkin.Api)
	mws.ZipkinURL = zipkinURL

	var zipkinTracer *zipkin.Tracer
	{
		if zipkinURL != "" {
			var (
				err      error

				hostPort = fmt.Sprintf("%s:%d", cfg.Com.Zipkin.Host, cfg.Com.Zipkin.Port)
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
			apphost := fmt.Sprintf("%s:%d", cfg.Com.Appdashot.Host, cfg.Com.Appdashot.Port)
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
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=%s&charset=%s&loc=",
		c.Com.DB.Username,
		c.Com.DB.Password,
		c.Com.DB.Host,
		c.Com.DB.Port,
		c.Com.DB.Database,
		c.Com.DB.ParseTime,
		c.Com.DB.Charset,
		url.QueryEscape(c.App.Timezone),
		)
	c.Bb, err = InitGORM(c.Com.DB.Dialect, args)

	return c.Bb, err
}