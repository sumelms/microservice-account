package config

import (
	"github.com/pkg/errors"
	"github.com/sherifabdlnaby/configuro"
	"os"
)

type Config struct {
	Server struct {
		Http *Server
	}
	Database *Database
	Logger   Logger
}

type Database struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Server struct {
	Host string
}

type Logger struct {
	Level string
	Debug bool
}

func NewConfig() (*Config, error) {
	configPath := os.Getenv("SUMELMS_CONFIG_PATH")
	if configPath == "" {
		return nil, errors.New("SUMELMS_CONFIG_PATH can not be empty.")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Wrapf(err, "Config file does not exists in %s", configPath)
	}

	loader, err := configuro.NewConfig(
		configuro.WithLoadFromConfigFile(configPath, false),
		configuro.WithLoadFromEnvVars("SUMELMS_"))
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := loader.Load(config); err != nil {
		return nil, err
	}

	return config, nil
}
