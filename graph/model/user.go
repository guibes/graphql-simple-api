package model

import (
	"github.com/dgrijalva/jwt-go"
)

// User model
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

// UserClaims for jwt
type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
