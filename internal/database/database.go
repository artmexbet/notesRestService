package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"notesRestService/internal/models"
)

type Config struct {
	Host   string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port   string `yaml:"port" env:"PORT" env-default:"5432"`
	User   string `yaml:"user" env:"USER" env-default:"postgres"`
	DBName string `yaml:"dbname" env:"DBNAME" env-default:"postgres"`
}

type Database struct {
	cfg  *Config
	conn *pgx.Conn
}

func New(cfg *Config) (*Database, error) {
	conn, err := pgx.Connect(context.Background(),
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.DBName))
	if err != nil {
		return nil, err
	}
	return &Database{cfg: cfg, conn: conn}, nil
}

func (d *Database) AddUser(user models.UserJSON) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var id int
	err = d.conn.QueryRow(context.Background(),
		"INSERT INTO public.users (login, password) VALUES ($1, $2)",
		user.Login, hashedPassword).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Database) CheckUser(user models.UserJSON) (int, error) {
	var hashedPassword []byte
	var id int
	err := d.conn.QueryRow(context.Background(),
		"SELECT id, password FROM public.users WHERE login=$1",
		user.Login).Scan(&id, &hashedPassword)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password))
	if err != nil {
		return 0, nil
	}
	return id, nil
}
