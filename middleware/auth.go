package middleware

import (
	"github.com/AvineshTripathi/orch/config"
	"net/http"
)

// AuthMiddleware checks for a valid Authorization token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != config.AuthToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
