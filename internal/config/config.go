package config

import (
	"log"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string 					`yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string 			`yaml:"storage_path" env:"STORAGE" env-default:".storage/storage.db"`
	HttpServer HttpServer       `yaml:"http_server"`
}

type HttpServer struct {
	Adress string             	`yaml:"adress" env:"ADRESS" env-default:"localhost:8082"`
	Timeout time.Duration 	  	`yaml:"timeout" env:"TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration 	`yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
}



func MustLoad() *Config {
	configPath := "./config/config.yml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %v does not exist", configPath)
	}
	

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil { log.Fatal(err.Error()) }

	return &cfg
}




