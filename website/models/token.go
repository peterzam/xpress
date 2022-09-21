package models

import jwt "github.com/golang-jwt/jwt"

// Token struct declaration
type Token struct {
	Id     uint
	Name   string
	Role   string
	Active bool
	*jwt.StandardClaims
}
