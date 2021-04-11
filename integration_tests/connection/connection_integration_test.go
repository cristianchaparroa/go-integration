//+build integration

package main

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"sftp/core"
	"testing"
)

// The following test just checks that is possible to
// connect to sftp server using a real connection with a docker
// container using the credentials in the file .env
func TestIntegrationSFTPConnexion(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = core.NewSFTPConnexion()
    assert.Nil(t, err)
}
