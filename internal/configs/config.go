package configs

type Config struct {
	Port   int    `json:"port"`
	Debug  bool   `json:"debug"`
	JwtKey string `json:"jwtKey"`
}

var config Config

func GetPort() int {
	if config.Port <= 0 {
		return 8080
	}
	return config.Port
}

func GetJwtKey() string {
	if len(config.JwtKey) <= 0 {
		return "0db23e471804d9ecbc126cbeb8393b89"
	}
	return config.JwtKey
}
