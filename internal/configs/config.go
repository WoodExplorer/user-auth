package configs

import "time"

type Config struct {
	Port   int    `json:"port"`
	Debug  bool   `json:"debug"`
	JwtKey string `json:"jwtKey"`
}

var (
	config Config

	jwtKeyUniqSuffix string
)

func init() {
	jwtKeyUniqSuffix = time.Now().String()
}

func GetPort() int {
	if config.Port <= 0 {
		return 8080
	}
	return config.Port
}

func GetJwtKey() string {
	// use jwtKeyUniqSuffix to make sure token will be invalid after service restart,
	// so as to avoid the problem caused by non-persistent token blacklist
	if len(config.JwtKey) <= 0 {
		return "0db23e471" + jwtKeyUniqSuffix
	}
	return config.JwtKey + jwtKeyUniqSuffix
}
