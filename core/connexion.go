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


func NewSFTPConnexion()  (SFTPConnexion, error) {
	sshAddress := os.Getenv("SSH_ADDRESS")
	sshUserName := os.Getenv("SSH_USERNAME")
	sshPassword := os.Getenv("SSH_PASSWORD")

	fmt.Println(sshAddress)

	conf := &ssh.ClientConfig{
		User: sshUserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
	}

	client, err := ssh.Dial("tcp", sshAddress, conf)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	connexion :=  &sftpConnexion{
		client: client,
	}

	return connexion, nil
}

func (s *sftpConnexion) Close() error {
	return s.Close()
}

