package Repositories

import (
	"errors"
	"task/Domain"
)

type UserRepository interface {
	// GetUserByUsername retrieves a user from the repository by their username.
	// If the user is not found, it returns an error.
	GetUserByUsername(username string) (*Domain.User, error)

	CreateUser(user *Domain.User) error
}

type userRepository struct {
	users []Domain.User
}

// returns a new instance of the userRepository struct.
func NewUserRepository() UserRepository {
	return &userRepository{users: []Domain.User{}}
}

// retrieves a user from the repository by their username.
func (r *userRepository) GetUserByUsername(username string) (*Domain.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// CreateUser adds a new user to the repository.
func (r *userRepository) CreateUser(user *Domain.User) error {
	r.users = append(r.users, *user)
	return nil
}
