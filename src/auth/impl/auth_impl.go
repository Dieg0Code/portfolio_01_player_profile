package impl

import (
	"errors"
	"time"

	"github.com/dieg0code/player-profile/src/auth"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

type AuthImpl struct{}

func (j *AuthImpl) GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func (j *AuthImpl) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
}

func NewJWTAth() auth.AuthUtils {
	return &AuthImpl{}
}
