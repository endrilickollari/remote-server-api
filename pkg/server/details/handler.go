package details

import (
	"encoding/json"
	"net/http"
)

// ServerDetailsHandler handles requests for server details
func ServerDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Example server details
	serverDetails := map[string]string{
		"hostname": "example.com",
		"status":   "running",
		"uptime":   "48 hours",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serverDetails)
}
