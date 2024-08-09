package Infrastructure

import (
	"errors"
	"task/Domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTService contains methods to generate and validate JWT tokens
type JWTService interface {
	GenerateJWT(username string) (string, error)
	// ValidateToken validates the given token string and returns the claims if the token is valid
	ValidateToken(tokenString string) (*Domain.Claims, error)
}

// jwtService is a concrete implementation of JWTService interface
type jwtService struct {
	SecretKey       string
	TokenExpiration time.Duration
}

// NewJWTService creates a new instance of jwtService with the given secret key and token expiration time
func NewJWTService(secretKey string, expiration time.Duration) JWTService {
	return &jwtService{
		SecretKey:       secretKey,
		TokenExpiration: expiration,
	}
}

// GenerateJWT generates a JWT token with the given username and the token expiration time
func (j *jwtService) GenerateJWT(username string) (string, error) {
	// Set the expiration time for the token
	expirationTime := time.Now().Add(j.TokenExpiration)
	// Create a new Claims object with the username and expiration time
	claims := &Domain.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Create a new JWT token with the claims and the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.SecretKey))
}

// ValidateToken validates the given token string and returns the claims if the token is valid
func (j *jwtService) ValidateToken(tokenString string) (*Domain.Claims, error) {
	// Create a new Claims object
	claims := &Domain.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
