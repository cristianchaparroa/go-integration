package core

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

type SFTPConnexion interface {
	Close() error
}

type sftpConnexion struct {
	client *ssh.Client
}

func NewSFTPConnexion() (SFTPConnexion, error) {
	sshHost := os.Getenv("SSH_HOST")
	sshPort := os.Getenv("SSH_PORT")
	sshUserName := os.Getenv("SSH_USERNAME")
	sshPassword := os.Getenv("SSH_PASSWORD")
	address := fmt.Sprintf("%s:%s", sshHost, sshPort)

	conf := &ssh.ClientConfig{
		User:            sshUserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
	}

	client, err := ssh.Dial("tcp", address, conf)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	connexion := &sftpConnexion{
		client: client,
	}

	return connexion, nil
}

func (s *sftpConnexion) Close() error {
	return s.Close()
}
