package Usecases

import (
	"errors"

	"task/Domain"
	"task/Infrastructure"
	"task/Repositories"

	"github.com/stretchr/testify/mock"
)

// MockUserUseCase is a mock implementation of the UserUseCase interface
type MockUserUseCase struct {
	mock.Mock
}

// UserUseCase is the implementation of the UserUseCase interface
type UserUseCase struct {
	UserRepo        Repositories.UserRepository
	JWTService      Infrastructure.JWTService
	PasswordService Infrastructure.PasswordService
}

// Register creates a new user in the repository
func (m *MockUserUseCase) Register(user *Domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Login authenticates a user and returns a JWT token if successful
func (m *MockUserUseCase) Login(credentials *Domain.Credentials) (string, error) {
	args := m.Called(credentials)
	return args.String(0), args.Error(1)
}

// Register creates a new user in the repository
func (uc *UserUseCase) Register(user *Domain.User) error {
	// Check if the user already exists
	existingUser, _ := uc.UserRepo.GetUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Hash the user's password
	hashedPassword, err := uc.PasswordService.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Set the user's password to the hashed password
	user.Password = hashedPassword
	return uc.UserRepo.CreateUser(user)
}

// Login authenticates a user and returns a JWT token if successful
func (uc *UserUseCase) Login(credentials *Domain.Credentials) (string, error) {
	// Get the user from the repository based on the username
	user, err := uc.UserRepo.GetUserByUsername(credentials.Username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Check if the provided password matches the user's password
	if !uc.PasswordService.ComparePasswords(user.Password, credentials.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate a JWT token for the user
	return uc.JWTService.GenerateJWT(user.Username)
}
