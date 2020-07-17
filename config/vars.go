package config

type Vars struct {
	NAME string
	SERVICE  string
	VERSION string
	MODE  uint8
	TIMEZONE string
	ENV string
	JWT_SECRET string
	BASE_PATH string
	CONFIG_PATH string
	LANGUAGE_PATH string
	DB_DIALECT  string
	DB_HOST string
	DB_PORT uint16
	DB_DATABASE string
	DB_USER string
	DB_PASSWORD string
	GRPC_PORT uint16
	DEBUG_PORT uint16
	REST_PORT  uint16
	REST_BASE_URL string
	ZIPKIN_HOST string
	ZIPKIN_API string
	ZIPKIN_PORT uint16
	ZIPKIN_BRIDGE bool
	LIGHTSTEP_TOKEN string
	APPDASHOT_HOST string
	APPDASHOT_PORT uint16
}
