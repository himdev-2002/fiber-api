package structs

import "github.com/golang-jwt/jwt/v5"

type JWTUserClaims struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}
