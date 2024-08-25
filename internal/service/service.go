package service

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

type IJWTManager interface {
	Encode(claims map[string]interface{}) (string, error)
	GetJWTAuth() *jwtauth.JWTAuth
}

type IDatabase interface {
}

type Config struct {
}

type Service struct {
	cfg        Config
	jwtManager IJWTManager
	db         IDatabase
}

func New(cfg Config, jwtManager IJWTManager, db IDatabase) *Service {
	return &Service{
		cfg:        cfg,
		jwtManager: jwtManager,
		db:         db,
	}
}

func (s *Service) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement registration logic here
	}
}

func (s *Service) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) AddNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) GetNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
