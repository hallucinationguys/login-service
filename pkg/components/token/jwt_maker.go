package token

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (*jwtMaker, error) {
	return &jwtMaker{secretKey}, nil
}

func (maker *jwtMaker) GenerateToken(email string, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, role, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))

	return token, payload, err
}

func (maker *jwtMaker) ValidationToken(token string) (*Payload, error) {

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(maker.secretKey), nil
	})

	log.Print(jwtToken)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	log.Print(jwtToken.Claims.(*Payload))
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
