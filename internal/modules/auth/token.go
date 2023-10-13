package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager interface {
	GenerateToken(id string) (string, error)
	ValidateToken(token string) (string, error)
}

type tokenManager struct {
	secret        string
	tokenLifeSpan int
}

func NewTokenManager(secret string, tokenLifeSpan int) TokenManager {
	return &tokenManager{secret, tokenLifeSpan}
}

func (t *tokenManager) GenerateToken(id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(t.tokenLifeSpan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}

func (t *tokenManager) ValidateToken(tokenString string) (string, error) {
	split := strings.Split(tokenString, " ")

	if len(split) != 2 {
		return "", fmt.Errorf("invalid token")
	}

	s := split[1]

	tk, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := tk.Claims.(jwt.MapClaims)

	if ok && tk.Valid {
		uid := claims["id"].(string)

		return uid, nil
	}

	return "", nil
}
