package router

import (
	"encoding/json"
	"net/http"
)

func (r *Router) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Login    string `json:"login"`
			Password string `json:"password`
		}

		req := &Request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}

		if err := r.service.Register(req.Login, req.Password); err != nil {
			http.Error(w, err.Error, http.StatusInternalServerError)
		}
	}
}

func (r *Router) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
