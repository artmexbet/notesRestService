package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"notesRestService/internal/models"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
)

type IJWTManager interface {
	Encode(claims map[string]interface{}) (string, error)
	GetJWTAuth() *jwtauth.JWTAuth
}

type IDatabase interface {
	AddUser(user models.UserJSON) (int, error)
	CheckUser(login string, password string) (models.UserID, error)
	AddNote(note models.NoteJSON, userId int) (int, error)
	GetNotes(userId int) (string, error)
}

type ITextValidator interface {
	ValidateTexts(text ...string) ([][]models.SpellerJSON, error)
}

type Config struct {
}

type Service struct {
	cfg           *Config
	jwtManager    IJWTManager
	db            IDatabase
	textValidator ITextValidator
	validator     *validator.Validate
}

func New(cfg *Config, jwtManager IJWTManager, db IDatabase, textValidator ITextValidator) *Service {
	return &Service{
		cfg:           cfg,
		jwtManager:    jwtManager,
		db:            db,
		textValidator: textValidator,
		validator:     validator.New(),
	}
}

func (s *Service) Register(lopgin, password string) error {
	user := &models.NewUser()

	if err := user.Validate(); err != nil {
		return err
	}
}

/*
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

		if err := s.validator.Struct(userJson); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Cannot validate json", slog.String("err", err.Error()))
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
*/

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

		if err := s.validator.Struct(userJson); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Cannot validate json", slog.String("err", err.Error()))
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

		logger.Info("User logged in", slog.Int("id", id))

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(models.TokenJSON{Token: token}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot encode json", slog.String("err", err.Error()))
		}
	}
}

func (s *Service) AddNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := slog.With(slog.String("module", "Service.AddNote"))

		var note models.NoteJSON
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Cannot decode json", slog.String("err", err.Error()))
			return
		}

		if err := s.validator.Struct(note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Cannot validate json", slog.String("err", err.Error()))
			return
		}

		spellerJSON, err := s.textValidator.ValidateTexts(note.Title, note.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot validate text", slog.String("err", err.Error()))
			return
		}

		if len(spellerJSON[0]) != 0 || len(spellerJSON[1]) != 0 {
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(spellerJSON); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logger.Error("Cannot encode json", slog.String("err", err.Error()))
			}
			return
		}
		start := time.Now()
		_, claims, _ := jwtauth.FromContext(r.Context())
		userId, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "Cannot get id from token", http.StatusInternalServerError)
			logger.Error("Cannot get id from token")
			return
		}
		fmt.Println(time.Since(start))

		noteId, err := s.db.AddNote(note, int(userId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot add note", slog.String("err", err.Error()))
			return
		}
		fmt.Println(time.Since(start))

		logger.Info("Note added", slog.Int("id", noteId))

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]any{"id": noteId}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot encode json", slog.String("err", err.Error()))
		}
	}
}

func (s *Service) GetNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := slog.With(slog.String("module", "Service.GetNotes"))

		_, claims, _ := jwtauth.FromContext(r.Context())
		userId, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "Cannot get id from token", http.StatusInternalServerError)
			logger.Error("Cannot get id from token")
			return
		}

		notes, err := s.db.GetNotes(int(userId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot get notes", slog.String("err", err.Error()))
			return
		}

		_, err = w.Write([]byte(notes))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Cannot write notes", slog.String("err", err.Error()))
		}
	}
}
