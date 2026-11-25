package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"pochistyam/internal/app"

	"golang.org/x/crypto/bcrypt"
)

// Reg Регистрирует нового пользователя
// @Summary      Регистрация пользователя
// @Description  Регистрирует нового пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "Данные комнаты"
// @Success      201 {object} AuthResponse
// @Failure      500 {string} string "StatusInternalServerError"
// @Router       /register [post]
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
		token, err := GenerateToken(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(AuthResponse{Token: token, UserID: userId, Name: req.Name})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Auth Аутентифицирует существующего пользователя
// @Summary      Аутентификация пользователя
// @Description  Аутентифицирует существующего пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body AuthRequest true "Данные комнаты"
// @Success      201 {object} AuthResponse
// @Failure      500 {string} string "StatusInternalServerError"
// @Router       /auth [post]
func Auth(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var userId, dbHash, name string
		err := application.DB.Pool.QueryRow(r.Context(), `SELECT id, password_hash, name FROM users WHERE email = $1`, req.Email).Scan(&userId, &dbHash, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(dbHash), []byte(req.Password)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := GenerateToken(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(AuthResponse{Token: token, UserID: userId, Name: name}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
