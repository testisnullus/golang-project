package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserID uint64 `json:"AccountID"`
}

const (
	AuthHeader = "Authorization"
)

var (
	ErrBadJwt = errors.New("JWT token invalid")
)

func NewJWT(userID uint64, hmacSecret []byte, tokenValidityTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenValidityTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(bearerToken string) (*Claims, error) {
	headerBearer := strings.SplitN(bearerToken, " ", 2)
	tokenString := headerBearer[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &Claims{})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %s", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("token claims are not in the expected format")
	}
	return claims, nil
}
