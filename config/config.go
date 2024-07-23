package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env    string     `yaml:"env" env-default:"local"`
	DB     Storage    `yaml:"storage"`
	Server HTTPServer `yaml:"http-server"`
}

type HTTPServer struct {
	Addr        string        `yaml:"address" env-default:"localhost:5000"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeOut time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}

func MustLoad() *Config {
	configPath := "/app/config/config.yaml"

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	return &cfg
}
