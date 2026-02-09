package config

type ApiConfig struct {
	Secret    string
	JWTIssuer string
	Domain    string
	Platform  string
}

func (cfg *ApiConfig) IsDebug() bool {
	if cfg.Platform == "debug" {
		return true
	}
	return false
}
