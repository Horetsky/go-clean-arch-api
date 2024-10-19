package services

import (
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateJWTToken(session dto.JWTSession) (string, error)
	GenerateJWTSession(user *entities.User) dto.JWTSession
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

func (s *jwtService) GenerateJWTToken(session dto.JWTSession) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	return token.SignedString([]byte("secret"))
}

func (s *jwtService) GenerateJWTSession(user *entities.User) dto.JWTSession {
	jwtUser := dto.JWTUser{
		ID:            user.ID,
		Email:         user.Email,
		Picture:       user.Picture,
		EmailVerified: user.EmailVerified,
	}

	if user.Talent != nil {
		jwtUser.TalentID = user.Talent.ID
	}

	if user.Recruiter != nil {
		jwtUser.RecruiterID = user.Recruiter.ID
	}

	return dto.JWTSession{
		User: jwtUser,
	}
}
