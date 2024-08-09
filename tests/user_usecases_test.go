package tests

import (
	"task/Domain"
	"task/Infrastructure"
	"task/Repositories"
	"task/Usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserUseCase_RegisterAndLogin(t *testing.T) {
	userRepo := Repositories.NewUserRepository()
	jwtService := Infrastructure.NewJWTService("test-secret", time.Minute*1)
	passwordService := Infrastructure.NewPasswordService()

	userUseCase := Usecases.UserUseCase{
		UserRepo:        userRepo,
		JWTService:      jwtService,
		PasswordService: passwordService,
	}

	// Test Register
	user := &Domain.User{
		Username: "testuser",
		Password: "securepassword",
		Role:     "user",
	}
	err := userUseCase.Register(user)
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)

	// Test Login with correct credentials
	credentials := Domain.Credentials{
		Username: "testuser",
		Password: "securepassword",
	}
	token, err := userUseCase.Login(&credentials)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test Login with incorrect credentials
	credentials.Password = "wrongpassword"
	token, err = userUseCase.Login(&credentials)
	assert.Error(t, err)
	assert.Empty(t, token)
}
