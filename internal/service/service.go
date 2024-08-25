package service

import (
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
	"net/http"
	"notesRestService/internal/models"
)

type IJWTManager interface {
	Encode(claims map[string]interface{}) (string, error)
	GetJWTAuth() *jwtauth.JWTAuth
}

type IDatabase interface {
	AddUser(user models.UserJSON) (int, error)
	CheckUser(user models.UserJSON) (int, error)
}

type Config struct {
}

type Service struct {
	cfg        *Config
	jwtManager IJWTManager
	db         IDatabase
}

func New(cfg *Config, jwtManager IJWTManager, db IDatabase) *Service {
	return &Service{
		cfg:        cfg,
		jwtManager: jwtManager,
		db:         db,
	}
}

func (s *Service) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := slog.With(slog.String("module", "Service.Register"))

		var userJson models.UserJSON
		err := json.NewDecoder(r.Body).Decode(&userJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Cannot decode json", slog.String("err", err.Error()))
			return
		}

		id, err := s.db.AddUser(userJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot add user", slog.String("err", err.Error()))
			return
		}
		logger.Info("User added", slog.Int("id", id))

		token, err := s.jwtManager.Encode(map[string]interface{}{"id": id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot encode token", slog.String("err", err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(models.TokenJSON{Token: token}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot encode json", slog.String("err", err.Error()))
			return
		}
	}
}

func (s *Service) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := slog.With(slog.String("module", "Service.Login"))

		var userJson models.UserJSON
		err := json.NewDecoder(r.Body).Decode(&userJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Cannot decode json", slog.String("err", err.Error()))
			return
		}

		id, err := s.db.CheckUser(userJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot check user", slog.String("err", err.Error()))
			return
		}

		if id == 0 {
			http.Error(w, "User not found", http.StatusNotFound)
			logger.Error("User not found")
			return
		}

		token, err := s.jwtManager.Encode(map[string]interface{}{"id": id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot encode token", slog.String("err", err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(models.TokenJSON{Token: token}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot encode json", slog.String("err", err.Error()))
		}
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
