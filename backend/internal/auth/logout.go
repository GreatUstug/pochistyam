package auth

import (
	"net/http"
	"pochistyam/internal/app"
	"pochistyam/internal/cache"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// LogOut Пользователь выходит из текущей сессии
// @Summary      Дает пользователю выйти из текущей сессии
// @Description  Выход пользователя из текущей сессии
// @Tags         auth
// @Success      200 {string} string
// @Failure      403 {string} string "StatusForbidden"
// @Router       /logOut [post]
func LogOut(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			jti, ok := claims["jti"].(string)
			if !ok {
				http.Error(w, "Invalid JTI", http.StatusForbidden)
				return
			}
			exp, ok := claims["exp"].(int64)
			if !ok {
				http.Error(w, "Invalid exp", http.StatusForbidden)
				return
			}
			time := time.Until(time.Unix(exp, 0))
			if cache.RedisClient != nil && jti != "" {
				cache.RedisClient.Set(r.Context(), "blacklist"+jti, "revoked", time)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"message":"logged out"}`))
		if err != nil {
			return
		}
	}
}
