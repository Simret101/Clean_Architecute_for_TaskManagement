package tests

import (
	"net/http"
	"net/http/httptest"
	"task/Infrastructure"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// generates a valid JWT token and sends a request to a protected endpoint with the token
// It asserts that the response code is 200 OK
func TestAuthMiddleware_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtService := Infrastructure.NewJWTService("secret", time.Hour)
	validToken, _ := jwtService.GenerateJWT("testuser")

	r := gin.Default()
	r.Use(Infrastructure.AuthMiddleware(jwtService))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// generates a invalid JWT token and sends a request to a protected endpoint with the token
// It asserts that the response code is 401 Unauthorized
func TestAuthMiddleware_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtService := Infrastructure.NewJWTService("secret", time.Hour)

	r := gin.Default()
	r.Use(Infrastructure.AuthMiddleware(jwtService))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
