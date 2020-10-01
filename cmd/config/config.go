package config

import (
	"sync"
)

type Config struct {
	Port        string `yaml:"Port"`
	Environment string `yaml:"Environment"`
}

var (
	ymlConf Config
	once    sync.Once
)

func Get() Config {
	once.Do(func() {
		ymlConf = Config{}
		ReadFromYml(&ymlConf)
	})
	return ymlConf
}
