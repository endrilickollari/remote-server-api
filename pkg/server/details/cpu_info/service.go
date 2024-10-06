package cpu_info

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"remote-server-api/pkg/login"
	"remote-server-api/pkg/utils"
	"strings"
)

func GetCPUInfo(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &login.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return login.JwtKey, nil
	})

	if err != nil || !tkn.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Retrieve SSH session using the session ID from the token
	client, exists := login.GetSession(claims.SessionID)
	if !exists {
		http.Error(w, "Session expired or not found", http.StatusUnauthorized)
		return
	}

	// Run command to get CPU info
	cpuInfoOutput, err := utils.RunCommand(client, "cat /proc/cpuinfo")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get CPU info: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse the CPU info
	cpuInfos, err := ParseCPUInfo(cpuInfoOutput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse CPU info: %v", err), http.StatusInternalServerError)
		return
	}

	// Return CPU details as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cpuInfos)
}
