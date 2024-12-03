package server

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UID *int64 `json:"uid"`
	jwt.RegisteredClaims
}

func (uc *UserClaims) GenerateToken(expiredTime time.Time, secret []byte) (string, error) {
	uc.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "ServerLib",
		Subject:   "Token",
		ExpiresAt: jwt.NewNumericDate(expiredTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uc *UserClaims) ParseToken(token string, secret []byte) error {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return err
	}
	claims, ok := tokenClaims.Claims.(*UserClaims)
	if !ok || !tokenClaims.Valid {
		return errors.New("the token claims is invalid")
	}
	*uc = *claims
	return nil
}
