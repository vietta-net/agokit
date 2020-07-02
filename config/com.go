package config

type Com struct {
	DB struct {
		Host      string `yaml:"host"`
		Port      uint16 `yaml:"port"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Database  string `yaml:"database"`
		Charset   string `yaml:"charset"`
		ParseTime string `yaml:"parsetime"`
	}

	Grpc struct {
		Host      string `yaml:"host"`
		Port      uint16 `yaml:"port"`
	}

	Rest struct {
		Host      string `yaml:"host"`
		Port      uint16 `yaml:"port"`
		baseurl   string `yaml:"baseurl"`
	}

	Zipkin struct {
		Host   string `yaml:"host"`
		Port   uint16 `yaml:"port"`
		Api    string `yaml:"api"`
		Bridge bool   `yaml:"bridge"`
	}

	Lightstep struct {
		Token string `yaml:"token"`
	}

	Appdashot struct {
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
	}

	Debug struct {
		Path  string   `yaml:"path"`
		Port uint16    `yaml:"port"`

	}

	Nats struct{
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
	}

	Smtp struct{
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		encryption  string `yaml:"encryption"`
		authentication  bool `yaml:"authentication"`
	}
}
