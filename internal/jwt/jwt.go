package jwt

import "github.com/go-chi/jwtauth/v5"

// Config ...
type Config struct {
	SecretKey string `yaml:"secret_key" env:"SECRET_KEY" env-default:"secret-key"`
}

// Manager ...
type Manager struct {
	cfg     *Config
	jwtAuth *jwtauth.JWTAuth
}

// New ...
func New(cfg *Config) *Manager {
	return &Manager{
		cfg:     cfg,
		jwtAuth: jwtauth.New("HS256", []byte(cfg.SecretKey), nil),
	}
}

func (m *Manager) GetJWTAuth() *jwtauth.JWTAuth {
	return m.jwtAuth
}

func (m *Manager) Encode(claims map[string]interface{}) (string, error) {
	_, tokenString, err := m.jwtAuth.Encode(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
