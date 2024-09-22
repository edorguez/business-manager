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

func (w *JWTWrapper) GenerateToken(userId int64, userEmail string, roleId int64, companyId int64, rolId int64, planId int64) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":        userId,
			"email":     userEmail,
			"roleId":    roleId,
			"companyId": companyId,
			"rolId":     rolId,
			"planId":    planId,
			"exp":       time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			"iss":       w.Issuer,
		})

	tokenString, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (w *JWTWrapper) ValidateToken(signedToken string) error {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(w.SecretKey), nil
	})

	if err != nil {
		return err
	}

	tokenExpirationDate, _ := token.Claims.GetExpirationTime()

	fmt.Println(tokenExpirationDate)
	fmt.Println(tokenExpirationDate.Unix())
	fmt.Println(time.Now().Unix())

	if !token.Valid || tokenExpirationDate.Unix() < time.Now().Unix() {
		return fmt.Errorf("Invalid token")
	}

	return nil
}
