package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	Environment string         `yaml:"environment"`
	Server      ServerConfig   `yaml:"server"`
	Database    DatabaseConfig `yaml:"database"`
}

// Loads the configuration
func LoadConfig() *Config {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("config/config.yaml", cfg); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	return cfg
}
