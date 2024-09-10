package router

import (
	"fmt"
	"net/http"
	"notesRestService/internal/models"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

type IService interface {
	Register(user models.User) error
	Login(login string, password string) (models.UserID, error)
	AddNote() http.HandlerFunc
	GetNotes() http.HandlerFunc
}

type IJWTManager interface {
	Encode(claims map[string]interface{}) (string, error)
	GetJWTAuth() *jwtauth.JWTAuth
}

// Config ...
type Config struct {
	Addr           string        `yaml:"host" env:"HOST" env-default:""`
	Port           string        `yaml:"port" env:"PORT" env-default:"8080"`
	Timeout        time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"60s"`
	MaxRequestSize int64         `yaml:"max_request_size" env:"MAX_REQUEST_SIZE" env-default:"41943040"` // 5MB
}

// Router ...
type Router struct {
	cfg     *Config
	router  *chi.Mux
	service IService
	jwt     IJWTManager
}

// New creates new Router instance
func New(cfg *Config, service IService, jwtManager IJWTManager) *Router {
	r := &Router{
		cfg:     cfg,
		router:  chi.NewRouter(),
		service: service,
		jwt:     jwtManager,
	}

	r.router.Use(middleware.Logger)
	r.router.Use(middleware.Recoverer)
	r.router.Use(middleware.RequestID)
	r.router.Use(middleware.AllowContentType("application/json"))
	r.router.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.router.Use(middleware.Timeout(cfg.Timeout))
	r.router.Use(middleware.RequestSize(cfg.MaxRequestSize))
	r.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.router.Post("/register", r.handleRegister())
	r.router.Post("/login", r.service.Login())

	r.router.Group(func(r_ chi.Router) {
		r_.Use(jwtauth.Verifier(jwtManager.GetJWTAuth()))
		r_.Use(jwtauth.Authenticator(jwtManager.GetJWTAuth()))

		r_.Post("/notes", r.service.NoteCreate())
		r_.Get("/notes", r.service.NoteList())
	})

	return r
}

func (r *Router) Run() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", r.cfg.Addr, r.cfg.Port), r.router)
}
