package jwt_manager

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

func (w *JWTWrapper) GenerateToken(userId int, userEmail string, role string) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    userId,
			"email": userEmail,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
			"iss":   w.Issuer,
		})

	tokenString, err := token.SignedString(w.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (w *JWTWrapper) ValidateToken(singedToken string) error {
	token, err := jwt.Parse(singedToken, func(token *jwt.Token) (interface{}, error) {
		return w.SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
