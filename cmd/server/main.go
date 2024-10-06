package main

import (
	"log"
	"net/http" // Update with your module name
	"remote-server-api/pkg/login"
)

func main() {
	// Initialize the HTTP server
	http.HandleFunc("/login", login.LoginHandler) // Handle the login endpoint

	port := "8080" // Change this to your desired port
	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
