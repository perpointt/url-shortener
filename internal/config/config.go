package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" envDefault:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTPPServer  `yaml:"http_server"`
}

type HTPPServer struct {
	Address     string `yaml:"address" env-default:"localhost:8080"`
	Timeout     int64  `yaml:"timeout" env-default:"4000"`
	IdleTimeout int64  `yaml:"idle_timeout" env-default:"6000"`
	User        string `yaml:"user" env-required:"true"`
	Password    string `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return &cfg
}
