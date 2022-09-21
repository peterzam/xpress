package models

import jwt "github.com/golang-jwt/jwt"

// Token struct declaration
type Token struct {
	Id   uint
	Name string
	Role string
	*jwt.StandardClaims
}
