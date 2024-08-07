package server

import (
	"testing"

	"github.com/jfelipearaujo-org/ms-customer-management/internal/environment"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("Should return a new server", func(t *testing.T) {
		// Arrange
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port: 5000,
			},
			DbConfig: &environment.DatabaseConfig{
				Url: "postgres://host:1234",
			},
			CloudConfig: &environment.CloudConfig{},
		}

		// Act
		server := NewServer(config)

		// Assert
		assert.NotNil(t, server)
	})

	t.Run("Should return a new server with base endpoint", func(t *testing.T) {
		// Arrange
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port: 5000,
			},
			DbConfig: &environment.DatabaseConfig{
				Url: "postgres://host:1234",
			},
			CloudConfig: &environment.CloudConfig{
				BaseEndpoint: "http://localhost:5000",
			},
		}

		// Act
		server := NewServer(config)

		// Assert
		assert.NotNil(t, server)
	})

	t.Run("Should create a http server", func(t *testing.T) {
		// Arrange
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port: 5000,
			},
			DbConfig: &environment.DatabaseConfig{
				Url: "postgres://host:1234",
			},
			CloudConfig: &environment.CloudConfig{},
		}

		server := NewServer(config)

		// Act
		httpServer := server.GetHttpServer()

		// Assert
		assert.NotNil(t, httpServer)
		assert.Equal(t, ":5000", httpServer.Addr)
	})
}
