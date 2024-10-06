package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
)

// LogInfo logs an informational message
func LogInfo(message string) {
	log.Println("INFO:", message)
}

// LogError logs an error message
func LogError(err error) {
	log.Println("ERROR:", err)
}

func TokenValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// Extract the token from the "Bearer <token>" format
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil // return the secret key for validation
		})

		if err != nil || !tkn.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Call the next handler if everything is valid
		next.ServeHTTP(w, r)
	}
}
