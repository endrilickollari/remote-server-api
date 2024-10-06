package login

import (
	"encoding/json"
	"fmt"
	"net/http"
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
