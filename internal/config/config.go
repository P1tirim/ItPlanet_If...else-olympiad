package config

import (
	"api/pkg/database/postgres"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	configErrRead      = "Read config error"
	configErrUnmarshal = "Unmarshal YAML config error"
	configErrProcess   = "Process ENV config error"
)

// Config struct.
type Config struct {
	DB  postgres.Config `yaml:"db" env:"API_DB"`
	App Server          `yaml:"app" json:"app"`
}

type Server struct {
	MaxProc    string `yaml:"maxProc" json:"maxProc" envconfig:"API_SERVER_MAXPROC"`
	Listenport string `yaml:"listenport" json:"listenport" envconfig:"API_SERVER_LISTENPORT"`
}

// GetConfig Получение конфига из файла.
func GetConfig(path string) (conf *Config, err error) {
	if conf, err = getYAMLConfig(path); err == nil {
		log.Println("get conf from file")

		return conf, nil
	}

	if conf, err = getEnvConfig(); err == nil {
		log.Println("get conf from env")

		return conf, nil
	}

	return nil, err
}

// getYAMLConfig return config from yaml file.
func getYAMLConfig(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, configErrRead)
	}

	config := new(Config)

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, errors.Wrap(err, configErrUnmarshal)
	}

	return config, nil
}

// getEnvConfig return env config.
func getEnvConfig() (*Config, error) {
	var conf Config

	envs := map[string]interface{}{
		"api_server": &conf.App,
		"api_db":     &conf.DB,
	}

	for name, object := range envs {
		if err := envconfig.Process(name, object); err != nil {
			return nil, errors.Wrap(err, configErrProcess)
		}
	}

	return &conf, nil
}
