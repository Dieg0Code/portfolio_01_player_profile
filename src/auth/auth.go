package auth

import "github.com/golang-jwt/jwt/v5"

type AuthUtils interface {
	GenerateToken(userID uint, role string) (string, error)
	ParseToken(tokenString string) (*jwt.Token, error)
}
