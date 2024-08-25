package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"notesRestService/internal/models"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	User     string `yaml:"user" env:"USER" env-default:"postgres"`
	Password string `yaml:"password" env:"PASSWORD" env-default:"postgres"`
	DBName   string `yaml:"dbname" env:"DBNAME" env-default:"postgres"`
}

type Database struct {
	cfg  *Config
	conn *pgx.Conn
}

func New(cfg *Config) (*Database, error) {
	conn, err := pgx.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
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

	row := d.conn.QueryRow(context.Background(),
		"INSERT INTO public.users (login, password) VALUES ($1, $2) RETURNING id",
		user.Login, hashedPassword)

	var id int
	err = row.Scan(&id)
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

func (d *Database) AddNote(note models.NoteJSON, userId int) (int, error) {
	var id int
	err := d.conn.QueryRow(context.Background(),
		"INSERT INTO public.notes (title, description, user_id) VALUES ($1, $2, $3) RETURNING id",
		note.Title, note.Description, userId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Database) GetNotes(userId int) (string, error) {
	row := d.conn.QueryRow(context.Background(),
		`SELECT coalesce(json_agg(row_to_json(notes)), '[]') FROM (
			SELECT id, title, description FROM public.notes WHERE user_id=$1) notes`, userId)

	var notes string
	err := row.Scan(&notes)
	if err != nil {
		return "", err
	}

	return notes, nil
}
