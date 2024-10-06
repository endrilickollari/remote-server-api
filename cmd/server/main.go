package main

import (
	"log"
	"net/http" // Update with your module name
	"remote-server-api/pkg/login"
	"remote-server-api/pkg/server/details"
	"remote-server-api/pkg/server/details/cpu_info"
)

func main() {
	// Initialize the HTTP server
	http.HandleFunc("/login", login.LoginHandler) // Handle the login endpoint
	http.HandleFunc("/server-details", login.TokenValidationMiddleware(details.ServerDetailsHandler))
	http.HandleFunc("/server-details/cpu-info", login.TokenValidationMiddleware(cpu_info.GetCPUInfo))

	port := "8080" // Change this to your desired port
	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
