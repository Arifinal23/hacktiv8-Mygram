package helper

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func NewJwt(id uint) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	expired, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))

	if err != nil {
		return "", err
	}

	jwtExpired := time.Now().Local().Add(time.Minute * time.Duration(expired))

	claims := jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: jwtExpired},
		Issuer:    "hacktiv8-mygram",
		Subject:   strconv.Itoa(int(id)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(secretKey)
}

func ParseJwt(tokenString string) (id uint, err error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	idInt, err := strconv.Atoi(claims["sub"].(string))

	if err != nil {
		return 0, err
	}

	return uint(idInt), nil
}
