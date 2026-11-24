package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"pochistyam/internal/app"

	"golang.org/x/crypto/bcrypt"
)

func Reg(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}(r.Body)
		if req.Name == "" || req.Email == "" || req.Password == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var userId string
		err = application.DB.Pool.QueryRow(r.Context(), `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id`, req.Name, req.Email, string(hash)).Scan(&userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(map[string]string{"userId": userId, "name": req.Name, "email": req.Email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Auth(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
