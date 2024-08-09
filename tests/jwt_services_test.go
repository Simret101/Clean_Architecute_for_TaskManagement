package tests

import (
	"testing"
	"time"

	"task/Infrastructure"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	secretKey := "test-secret"
	expiration := time.Minute * 1

	jwtService := Infrastructure.NewJWTService(secretKey, expiration)
	username := "testuser"

	// Test GenerateJWT
	tokenString, err := jwtService.GenerateJWT(username)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Test ValidateToken
	claims, err := jwtService.ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, username, claims.Username)
}

func TestValidateToken_InvalidToken(t *testing.T) {
	secretKey := "test-secret"
	expiration := time.Minute * 1

	jwtService := Infrastructure.NewJWTService(secretKey, expiration)

	invalidToken := "invalid.token.string"
	claims, err := jwtService.ValidateToken(invalidToken)
	assert.Error(t, err)
	assert.Nil(t, claims)
}
