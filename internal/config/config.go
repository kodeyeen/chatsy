package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string        `yaml:"env" env-required:"true"`
	Secret     string        `yaml:"secret" env-required:"true"`
	TokenTTL   time.Duration `yaml:"token_ttl" env-required:"true"`
	TicketTTL  time.Duration `yaml:"ticket_ttl" env-required:"true"`
	HTTPServer HTTPServer    `yaml:"http_server"`
	Cors       Cors          `yaml:"cors"`
	Database   Database      `yaml:"database"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Cors struct {
	AllowedOrigins []string `yaml:"allowed_origins" env-required:"true"`
	Debug          bool     `yaml:"debug" env-required:"true"`
}

type Database struct {
	Host     string `yaml:"host" env-required:"true"`
	Name     string `yaml:"name" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

func MustLoad() *Config {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	_, err := os.Stat(cfgPath)
	if os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", cfgPath)
	}

	var cfg Config
	err = cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
