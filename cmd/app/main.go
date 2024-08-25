package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"notesRestService/internal/database"
	"notesRestService/internal/jwt"
	"notesRestService/internal/logger/sl"
	"notesRestService/internal/router"
	"notesRestService/internal/service"
)

type Config struct {
	JWTConfig      *jwt.Config      `yaml:"jwt" env-prefix:"JWT_"`
	RouterConfig   *router.Config   `yaml:"router" env-prefix:"ROUTER_"`
	DatabaseConfig *database.Config `yaml:"db" env-prefix:"DB_"`
	ServiceConfig  *service.Config  `yaml:"service" env-prefix:"SERVICE_"`
}

// readConfig ...
func readConfig(filename string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filename, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	sl.SetupLogger("local")

	cfg, err := readConfig("./config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.New(cfg.DatabaseConfig)
	if err != nil {
		log.Fatalln(err)
	}

	jwtManager := jwt.New(cfg.JWTConfig)

	svc := service.New(cfg.ServiceConfig, jwtManager, db)

	r := router.New(cfg.RouterConfig, svc, jwtManager)
	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
