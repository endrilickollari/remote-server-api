package login

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// LoginHandler handles the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var sshLogin SSHLogin

	err := json.NewDecoder(r.Body).Decode(&sshLogin)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Connect to SSH
	err = ConnectToSSH(sshLogin.IP, sshLogin.Username, sshLogin.Port, sshLogin.Password)
	if err != nil {
		http.Error(w, "SSH connection failed", http.StatusUnauthorized)
		return
	}

	// Generate token
	token, err := GenerateToken(sshLogin.Username)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
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
