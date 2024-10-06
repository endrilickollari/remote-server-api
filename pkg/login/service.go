package login

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4" // For JWT generation
	"golang.org/x/crypto/ssh"      // For SSH connection
	"sync"
	"time"
)

// JWT secret key (replace with something more secure in production)
var JwtKey = []byte("your_secret_key")

// SSHSessions is a map that stores SSH client sessions
var SSHSessions = make(map[string]*ssh.Client)
var sessionMutex sync.RWMutex

// StoreSession saves the SSH client in the map
func StoreSession(sessionID string, client *ssh.Client) {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()
	SSHSessions[sessionID] = client
}

// GetSession retrieves the SSH client from the map
func GetSession(sessionID string) (*ssh.Client, bool) {
	sessionMutex.RLock()
	defer sessionMutex.RUnlock()
	client, exists := SSHSessions[sessionID]
	return client, exists
}

// RemoveSession deletes an SSH session (e.g., when the user logs out or session expires)
func RemoveSession(sessionID string) {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()
	delete(SSHSessions, sessionID)
}

func ConnectToSSH(ip, username, port, password string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	connectionString := fmt.Sprintf("%s:%s", ip, port)
	client, err := ssh.Dial("tcp", connectionString, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial SSH: %v", err)
	}

	// Return the SSH client for use in future calls
	return client, nil
}

func GenerateToken(username, sessionID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Set token expiration
	claims := &Claims{
		Username:  username,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
