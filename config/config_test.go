package config_test

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vietta-net/agokit/config"

)

func TestNew(t *testing.T) {
	os.Setenv("APP_TEST", "true")
	os.Setenv("APP_PREFIX", "APP")


	//NAME string
	var n = "NAME"

	t.Run(n, func(t *testing.T) {
		var v = "MyApp"
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.Name)
	})

	//SERVICE  string
	n = "SERVICE"
	t.Run(n, func(t *testing.T) {
		var v = "Masterdata"
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.Service)
	})
	//VERSION string
	n = "VERSION"
	t.Run(n, func(t *testing.T) {
		var v = "v.0.0.1"
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.Version)
	})
	//MODE  uint8
	n = "MODE"
	t.Run(n, func(t *testing.T) {
		var v = uint8(1)
		os.Setenv("APP_" + n, "1")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.Mode)
	})
	//TIMEZONE string
	n = "TIMEZONE"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.Timezone)
	})
	//ENV string
	n = "ENV"
	t.Run(n, func(t *testing.T) {
		var v = "dev"
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.Env)
	})
	//JWT_SECRET string
	n = "JWT_SECRET"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.Secret)
	})

	//BASE_PATH string
	n = "BASE_PATH"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.App.BasePath)
	})
	//CONFIG_PATH string
	n = "CONFIG_PATH"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Arg.ConfigPath)
	})
	//LANGUAGE_PATH string
	n = "LANGUAGE_PATH"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Arg.LanguagePath)
	})
	//DB_DIALECT  string
	n = "DB_DIALECT"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.DB.Dialect)
	})
	//DB_HOST string
	n = "DB_HOST"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.DB.Host)
	})
	//DB_PORT uint16
	n = "DB_PORT"
	t.Run(n, func(t *testing.T) {
		var v = uint16(33060)
		os.Setenv("APP_" + n, "33060")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.DB.Port)
	})
	//DB_DATABASE string
	n = "DB_DATABASE"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.DB.Database)
	})
	//DB_USER string
	n = "DB_USER"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.DB.Username)
	})
	//DB_PASSWORD string
	n = "DB_PASSWORD"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.DB.Password)
	})
	//GRPC_PORT uint16
	n = "GRPC_PORT"
	t.Run(n, func(t *testing.T) {
		var v = uint16(6400)
		os.Setenv("APP_" + n, "6400")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Grpc.Port)
	})
	//DEBUG_PORT uint16
	n = "DEBUG_PORT"
	t.Run(n, func(t *testing.T) {
		var v = uint16(6400)
		os.Setenv("APP_" + n, "6400")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Debug.Port)
	})
	//REST_PORT  uint16
	n = "REST_PORT"
	t.Run(n, func(t *testing.T) {
		var v = uint16(6400)
		os.Setenv("APP_" + n, "6400")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Rest.Port)
	})
	//REST_BASE_URL string
	n = "REST_BASE_URL"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Rest.BaseUrl)
	})
	//ZIPKIN_HOST string
	n = "ZIPKIN_HOST"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Zipkin.Host)
	})
	//ZIPKIN_API string
	n = "ZIPKIN_API"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Zipkin.Api)
	})
	//ZIPKIN_PORT uint16
	n = "ZIPKIN_PORT"
	t.Run(n, func(t *testing.T) {
		var v = uint16(64000)
		os.Setenv("APP_" + n, "64000")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Zipkin.Port)
		t.Log(basicConfig.Com.Zipkin.Port)
	})
	//ZIPKIN_BRIDGE bool
	n = "ZIPKIN_BRIDGE"
	t.Run(n, func(t *testing.T) {
		var v = false
		os.Setenv("APP_" + n, "false")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Zipkin.Bridge)
		t.Log(basicConfig.Com.Zipkin.Bridge)
	})
	//LIGHTSTEP_TOKEN string
	n = "LIGHTSTEP_TOKEN"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Lightstep.Token)
	})
	//APPDASHOT_HOST string
	n = "APPDASHOT_HOST"
	t.Run(n, func(t *testing.T) {
		var v = n
		os.Setenv("APP_" + n, v)
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Appdashot.Host)
	})
	//APPDASHOT_PORT uint16
	n = "APPDASHOT_PORT"
	t.Run(n, func(t *testing.T) {
		var v = uint16(64000)
		os.Setenv("APP_" + n, "64000")
		cfg, err := config.New("..")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		basicConfig :=  cfg.(*config.BasicConfig)
		assert.Equal(t, v, basicConfig.Com.Appdashot.Port)
	})


}