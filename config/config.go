package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config - config struct
type Config struct {

	// LogLevel            string `envconfig:"LOG_LEVEL"`

	HTTPAddr string `envconfig:"HTTP_ADDR"`
	LogPath  string `envconfig:"LOG_FILE_PATH"`

	// database settings:
	DB DBConfig
}

// DBConfig is config for database
type DBConfig struct {
	Username string `envconfig:"DB_USERNAME"`
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	DBName   string `envconfig:"DB_NAME"`
	SSLMode  string `envconfig:"DB_SSL_MODE"`
	Password string `envconfig:"DB_PASSWORD"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
