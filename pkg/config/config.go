package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sherifabdlnaby/configuro"
)

type Config struct {
	Service string // Service name
	Server  struct {
		Http *Server
		Grpc *Server
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

func NewConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Wrapf(err, "Config file does not exists in %s", configPath)
	}

	loader, err := configuro.NewConfig(
		configuro.WithLoadFromConfigFile(configPath, false),
		configuro.WithLoadFromEnvVars("SUMELMS"))
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := loader.Load(config); err != nil {
		return nil, err
	}

	return config, nil
}
