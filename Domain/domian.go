package Domain

import (
	"github.com/golang-jwt/jwt/v4"
)

// Task represents a task entity
type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     string
	Status      string
}

// User represents a user entity
type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

// Credentials represents user login credentials
type Credentials struct {
	Username string
	Password string
}

// Claims represents the JWT claims embedded in the token
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
