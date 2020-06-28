package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server  server
	MySQL   MySQL
	Logging logging
}

type server struct {
	ServerAddress      string
	ServerRetries      int
	ServerRetryTimeout int
}

type MySQL struct {
	Host     string
	Username string
	Password string
	Database string
	Timezone string
}

type logging struct {
	ApiLogPath string
	LogLevel   string
}

var C *Config

func loadConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("Config File Does Not Exist")
	} else if err != nil {
		return nil, err
	}

	var conf Config
	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func InitializeConfig(configPath string) {
	temp, err := loadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to Initialize Configuration: %s", err.Error()))
	}
	C = temp
}
