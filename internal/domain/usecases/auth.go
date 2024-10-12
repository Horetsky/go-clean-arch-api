package usecases

import (
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"
	errs "seeker/internal/domain/errors"
	"seeker/internal/domain/services"
	"seeker/internal/domain/storages"
	"seeker/internal/types"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(input dto.RegisterUserInput) (types.JWTTokenResponse, types.JWTSession, error)
	Login(input dto.LoginUserInput) (types.JWTTokenResponse, types.JWTSession, error)
	GenerateSession(user *entities.User) (types.JWTTokenResponse, types.JWTSession, error)
}

type authUsecase struct {
	userRepository storages.UserStorage
	jwtService     services.JWTService
}

func NewAuthUsecase(userRepository storages.UserStorage) AuthUsecase {
	jwtService := services.NewJWTService()
	return &authUsecase{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (u *authUsecase) Register(input dto.RegisterUserInput) (types.JWTTokenResponse, types.JWTSession, error) {
	_, err := u.userRepository.GetByEmail(input.Email)
	var tokens types.JWTTokenResponse
	var session types.JWTSession

	// e == nil means that user exists
	if err == nil {
		return tokens, session, errs.ErrUserAlreadyExists
	}

	newUser := &entities.User{
		Email: input.Email,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return tokens, session, errs.ErrFailedToCreateUser
	}

	newUser.Password = string(hashedPassword)

	err = u.userRepository.CreateOne(newUser)

	if err != nil {
		return tokens, session, err
	}

	tokens, session, err = u.GenerateSession(newUser)

	if err != nil {
		return tokens, session, err
	}

	return tokens, session, nil
}

func (u *authUsecase) Login(input dto.LoginUserInput) (types.JWTTokenResponse, types.JWTSession, error) {
	dbUser, err := u.userRepository.GetByEmail(input.Email)
	var tokens types.JWTTokenResponse
	var session types.JWTSession

	if err != nil {
		return tokens, session, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password))

	if err != nil {
		return tokens, session, errs.ErrInvalidPassword
	}

	tokens, session, err = u.GenerateSession(&dbUser)

	if err != nil {
		return tokens, session, err
	}

	return tokens, session, nil
}

func (u *authUsecase) GenerateSession(user *entities.User) (types.JWTTokenResponse, types.JWTSession, error) {
	var tokens types.JWTTokenResponse
	var session types.JWTSession

	session = u.jwtService.GenerateJWTSession(user)

	accessToken, err := u.jwtService.GenerateJWTToken(session)
	refreshToken, err := u.jwtService.GenerateJWTToken(session)

	if err != nil {
		return tokens, session, errs.ErrFailedToCreateSession
	}

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken

	return tokens, session, nil
}
