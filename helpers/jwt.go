package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userID string, role string) (string, error) {
	jc := jwtClaim{userID, role, jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * 2400).Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc)

	// TODO: parse jwt secret through env
	tokenString, err := token.SignedString([]byte("abcd"))
	fmt.Println(err)

	return tokenString, nil
}
