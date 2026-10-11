package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secret []byte
}

func NewJwtService(secret string) *JwtService {
	return &JwtService{
		secret: []byte(secret),
	}
}

func (s *JwtService) GerarToken(idPessoa string) (string, error) {
	agora := time.Now()

	claims := jwt.RegisteredClaims{
		Issuer:    "rozae",
		Subject:   idPessoa,
		IssuedAt:  jwt.NewNumericDate(agora),
		ExpiresAt: jwt.NewNumericDate(agora.Add(time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	retorno, err := token.SignedString(s.secret)
	return retorno, err
}

func (s *JwtService) ValidarToken(tokenString string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (any, error) {
			return s.secret, nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer("rozae"),
	)

	if err != nil || !token.Valid {
		return nil, errors.New("Token inválido")
	}

	if claims.Subject == "" {
		return nil, errors.New("Pessoa não identificada")
	}

	return claims, nil
}
