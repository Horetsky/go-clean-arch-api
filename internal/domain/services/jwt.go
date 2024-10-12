package services

import (
	"seeker/internal/domain/entities"
	"seeker/internal/types"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateJWTToken(session types.JWTSession) (string, error)
	GenerateJWTSession(user *entities.User) types.JWTSession
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

func (s *jwtService) GenerateJWTToken(session types.JWTSession) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	return token.SignedString([]byte("secret"))
}

func (s *jwtService) GenerateJWTSession(user *entities.User) types.JWTSession {
	return types.JWTSession{
		User: types.JWTUser{
			ID:      user.ID,
			Email:   user.Email,
			Picture: user.Picture,
		},
	}
}
