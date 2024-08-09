package tests

import (
	"task/Infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordService(t *testing.T) {
	passwordService := Infrastructure.NewPasswordService()
	password := "securepassword"

	// Test HashPassword
	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Test ComparePasswords
	isValid := passwordService.ComparePasswords(hashedPassword, password)
	assert.True(t, isValid)

	// Test ComparePasswords with wrong password
	isValid = passwordService.ComparePasswords(hashedPassword, "wrongpassword")
	assert.False(t, isValid)
}
