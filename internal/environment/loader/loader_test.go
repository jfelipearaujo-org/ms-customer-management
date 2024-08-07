package loader

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/jfelipearaujo-org/ms-customer-management/internal/environment"
	"github.com/stretchr/testify/assert"
)

func cleanEnv() {
	envs := []string{
		"API_PORT",
		"API_ENV_NAME",
		"API_VERSION",
		"DB_URL",
		"DB_URL_SECRET_NAME",
		"AWS_BASE_ENDPOINT",
	}

	for _, env := range envs {
		os.Unsetenv(env)
	}
}

func TestGetEnvironment(t *testing.T) {
	t.Run("Should load environment variables", func(t *testing.T) {
		// Arrange
		envs := []struct {
			name  string
			value string
		}{
			{"API_PORT", "5000"},
			{"API_ENV_NAME", "development"},
			{"API_VERSION", "v1"},
			{"DB_URL", "db://host:1234"},
			{"DB_URL_SECRET_NAME", "db-secret-url"},
			{"AWS_BASE_ENDPOINT", "http://localhost:4566"},
		}

		for _, env := range envs {
			t.Setenv(env.name, env.value)
		}

		expected := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port:       5000,
				EnvName:    "development",
				ApiVersion: "v1",
			},
			DbConfig: &environment.DatabaseConfig{
				Url:           "db://host:1234",
				UrlSecretName: "db-secret-url",
			},
			CloudConfig: &environment.CloudConfig{
				BaseEndpoint: "http://localhost:4566",
			},
		}

		// Act
		env, err := NewLoader().GetEnvironment(context.Background())

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, env)
		assert.Equal(t, expected, env)
	})

	t.Run("Should return error if a required variable is not set", func(t *testing.T) {
		// Arrange
		envs := []struct {
			name  string
			value string
		}{
			{"API_PORT", "5000"},
			{"API_ENV_NAME", "development"},
			{"API_VERSION", "v1"},
			{"DB_URL", "db://host:1234"},
			{"AWS_BASE_ENDPOINT", "http://localhost:4566"},
		}

		for _, env := range envs {
			t.Setenv(env.name, env.value)
		}

		// Act
		env, err := NewLoader().GetEnvironment(context.Background())

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
	})
}

func TestGetEnvironmentFromFile(t *testing.T) {
	t.Run("Should load environment variables from file", func(t *testing.T) {
		// Arrange
		cleanEnv()

		expected := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port:       5000,
				EnvName:    "development",
				ApiVersion: "v1",
			},
			DbConfig: &environment.DatabaseConfig{
				Url:           "db://host:1234",
				UrlSecretName: "db-secret-url",
			},
			CloudConfig: &environment.CloudConfig{
				BaseEndpoint: "http://localhost:4566",
			},
		}

		// Act
		env, err := NewLoader().GetEnvironmentFromFile(context.Background(), filepath.Join("testdata", "test.env"))

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, env)
		assert.Equal(t, expected, env)
	})

	t.Run("Should return error if a required variable is not set", func(t *testing.T) {
		// Arrange
		cleanEnv()

		// Act
		env, err := NewLoader().GetEnvironmentFromFile(context.Background(), filepath.Join("testdata", "test_err.env"))

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
	})

	t.Run("Should return error try to load from an invalid file", func(t *testing.T) {
		// Arrange
		cleanEnv()

		// Act
		env, err := NewLoader().GetEnvironmentFromFile(context.Background(), filepath.Join("testdata", "non_exists.env"))

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
	})
}
