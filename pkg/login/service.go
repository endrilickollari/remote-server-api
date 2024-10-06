package login

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4" // For JWT generation
	"golang.org/x/crypto/ssh"      // For SSH connection
	"time"
)

// JWT secret key (replace with something more secure in production)
var jwtKey = []byte("your_secret_key")

func ConnectToSSH(ip, username, port, password string) error {
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
		return err
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	return nil
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
