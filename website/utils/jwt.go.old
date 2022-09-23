package utils

import (
	"Xpress/models"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Token{
		Id:     user.Id,
		Name:   user.Name,
		Role:   user.Role,
		Active: user.Active,
	})
	tokenString, err := token.SignedString([]byte(GetEnv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
