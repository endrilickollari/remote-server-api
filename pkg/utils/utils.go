package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func LogInfo(message string) {
	log.Println("INFO:", message)
}

func LogError(err error) {
	log.Println("ERROR:", err)
}

func RunCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	if err := session.Run(command); err != nil {
		return "", fmt.Errorf("failed to run command '%s': %v", command, err)
	}

	return stdoutBuf.String(), nil
}
