package helpers

import (
	"strings"
	"time"

	"github.com/brianvoe/sjwt"
	"github.com/burntcarrot/pm/errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtClaim struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userName string, role string) (string, error) {
	jc := jwtClaim{userName, role, jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * 2400).Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc)

	// TODO: parse jwt secret through env
	tokenString, err := token.SignedString([]byte("abcd"))
	if err != nil {
		return "", errors.ErrInternalServerError
	}

	return tokenString, nil
}

func UserRoleValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		role, err := ExtractJWTPayloadRole(e)
		if err != nil {
			return echo.ErrUnauthorized
		}
		if role == "user" {
			return next(e)
		}
		return echo.ErrUnauthorized
	}
}

func ExtractJWTPayloadRole(c echo.Context) (string, error) {
	header := c.Request().Header.Clone().Get("Authorization")
	if header == "" {
		return "", errors.ErrInvalidCredentials
	}

	token := strings.Split(header, "Bearer ")[1]
	if token == "" {
		return "", errors.ErrInvalidCredentials
	}

	claims, err := sjwt.Parse(token)
	if err != nil {
		return "", err
	}

	if claims["role"] == nil {
		return "", errors.ErrInvalidCredentials
	}

	return claims["role"].(string), nil
}
