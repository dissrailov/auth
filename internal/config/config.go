package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	JWT        `yaml:"jwt"`
	HTTPServer `yaml:"http_server"`
	Database   `yaml:"database"`
}

type HTTPServer struct {
	Addr        string        `yaml:"addr" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}

type JWT struct {
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env-default:"24h"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env-default:"24h"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}
	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	return &config
}
