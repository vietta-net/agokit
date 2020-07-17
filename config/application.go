package config

type Application struct {
	Name string 		`yaml:"name" json:"name"`
	Service string 		`yaml:"service" json:"service"`
	Version string 		`yaml:"version" json:"version"`
	Env string			`yaml:"env" json:"env"`
	Debug bool 			`yaml:"debug" json:"debug"`
	BasePath string 	`yaml:"basepath" json:"basepath"`
	Mode uint8			`yaml:"mode" json:"mode"` //1 - Real, 0 - Test
	Language string		`yaml:"language" json:"language"`
	Timezone string 	`yaml:"timezone" json:"timezone"`
	Secret string		`yaml:"secret" json:"secret"` //JWT Secret
	AcceptLanguage string `yaml:"acceptLanguage" json:"acceptLanguage"`
}

