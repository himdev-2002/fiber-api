package helpers

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"tde/fiber-api/core/models"
	"tde/fiber-api/core/services"
	"tde/fiber-api/core/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lithammer/shortuuid"
)

func GenerateToken(user *models.User) (string, string, error) {
	var err error
	var token string
	var refreshToken string

	expiredHour := 1
	expiredHour, err = strconv.Atoi(os.Getenv("JWT_EXPIRED_HOUR"))
	if err != nil {
		if log, err2 := services.ErrorLog(); err2 == nil {
			msg := err.Error()
			log.Log().Msgf(msg)
		}
	} else {
		uuid := shortuuid.New()
		// fmt.Println(uuid)
		claims := structs.JWTUserClaims{
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  *user.LastName,
			RegisteredClaims: jwt.RegisteredClaims{
				// A usual scenario is to set the expiration time relative to the current time
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredHour) * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    os.Getenv("JWT_ISSUER"),
				Subject:   user.Username,
				ID:        uuid,
				// Audience:  []string{"somebody_else"},
			},
		}
		// fmt.Println(claims)

		refreshClaims := structs.JWTUserClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				// A usual scenario is to set the expiration time relative to the current time
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredHour*2) * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    os.Getenv("JWT_ISSUER"),
				Subject:   user.Username,
				ID:        uuid,
				// Audience:  []string{"somebody_else"},
			},
		}
		// fmt.Println(refreshClaims)

		token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))
		// fmt.Println(token)
		if err != nil {
			if log, err2 := services.ErrorLog(); err2 == nil {
				msg := err.Error()
				log.Log().Msgf(msg)
			}
		} else {
			refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("SECRET_KEY")))
		}
	}

	// fmt.Println(token, refreshToken, err)
	return token, refreshToken, err
}

func ValidateToken(tokenString *string) (*structs.JWTUserClaims, error) {
	token, err := jwt.ParseWithClaims(*tokenString, &structs.JWTUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		if log, err2 := services.ErrorLog(); err2 == nil {
			msg := err.Error()
			log.Log().Msgf(msg)
		}
		return nil, err
	} else if claims, ok := token.Claims.(*structs.JWTUserClaims); ok {
		if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
			if log, err := services.InfoLog(); err == nil {
				log.Info().Msgf("token is expired")
			}
			return nil, fmt.Errorf("token is expired")
		} else {
			return claims, nil
		}
	} else {
		if log, err := services.ErrorLog(); err == nil {
			log.Fatal().Msgf("unknown claims type, cannot proceed")
		}
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}

}

func CheckPublicRouter(path *string, publicRoute *[]string) bool {
	ret := false

	if slices.Contains(*publicRoute, *path) {
		ret = true
	}

	return ret
}
