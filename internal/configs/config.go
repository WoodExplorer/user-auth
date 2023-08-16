package configs

type Config struct {
	Port  int  `json:"port"`
	Debug bool `json:"debug"`
}

var config Config

func GetPort() int {
	if config.Port <= 0 {
		return 8080
	}
	return config.Port
}
