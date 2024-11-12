package jwt

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type (
	JWT struct {
		key []byte
	}

	MyCustomClaims struct {
		UserID int
		jwt.RegisteredClaims
	}
)

func NewJwt(conf *viper.Viper) *JWT {
	return &JWT{key: []byte(conf.GetString("security.jwt.key"))}
}

func (j *JWT) GenToken(userID int, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	if strings.TrimSpace(tokenString) == "" {
		return nil, errors.New("token is empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
