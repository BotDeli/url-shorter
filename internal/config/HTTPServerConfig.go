package config

import "time"

type HTTPServerConfig struct {
	Address           string        `yaml:"address" env-default:"localhost"`
	Port              string        `yaml:"port" env-default:"8080"`
	Protocol          string        `yaml:"protocol" env-default:"http://"`
	ReadHeaderTimeout time.Duration `yaml:"read-header-timeout" env-default:"4"`
	IdleTimeout       time.Duration `yaml:"idle-timeout" env-default:"60"`
}

func (cfg *HTTPServerConfig) GetHomeAddress() string {
	return cfg.Address + ":" + cfg.Port
}
