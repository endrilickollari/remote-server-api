package utils

import (
	"log"
)

// LogInfo logs an informational message
func LogInfo(message string) {
	log.Println("INFO:", message)
}

// LogError logs an error message
func LogError(err error) {
	log.Println("ERROR:", err)
}
