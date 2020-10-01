package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

/*
	The config package can load application configuration from yaml files inside config directory.
	The basic config includes server port and go_environment. GoEnvironment and Port has default values
	that can be override with environment variables.

	This is useful to provide configuration on containers.

	The default values are:
	- Port: 8080
	- GoEnvironment: develop (this means that additional config will be loaded from config/develop.yml)
*/

const (
	defaultPort      = "8080"
	defaultGoEnv     = "develop"
	filePathFormat   = "%s/config/%s.yml"
	envGoEnvironment = "GO_ENVIRONMENT"
	envPort          = "PORT"
)

func ReadFromYml(config *Config) {
	env, port := readEnv()
	config.Port = port
	config.Environment = env

	yamlFile, err := ioutil.ReadFile(getFileName(env))
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func getFileName(env string) string {
	basePath, _ := os.Getwd()
	filePath := fmt.Sprintf(filePathFormat, basePath, env)
	return filePath
}

func readEnv() (env, port string) {
	env = os.Getenv(envGoEnvironment)
	port = os.Getenv(envPort)

	if env == "" {
		env = defaultGoEnv
	}

	if port == "" {
		port = defaultPort
	}

	return
}
