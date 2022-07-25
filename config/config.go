package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type Config struct {
	Protocol string
	Host     string
	Port     string
}

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(basepath)
	if os.Getenv("ENV") == "dev" {
		viper.SetConfigName("config-dev.yaml")
	} else {
		viper.SetConfigName("config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// ParseConfig Parse config file
func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
