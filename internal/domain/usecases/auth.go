package usecases

import (
	"errors"
	"log"
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"
	errs "seeker/internal/domain/errors"
	"seeker/internal/domain/repositories"
	"seeker/internal/domain/services"
	"seeker/internal/types"

	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(input dto.RegisterUserInput) (types.JWTTokenResponse, types.JWTSession, error)
	Login(input dto.LoginUserInput) (types.JWTTokenResponse, types.JWTSession, error)
	GenerateSession(user *entities.User) (types.JWTTokenResponse, types.JWTSession, error)
	VerifyEmail(email string) (types.JWTTokenResponse, types.JWTSession, error)
}

type authUsecase struct {
	userRepository repositories.UserRepository
	jwtService     services.JWTService
	emailService   services.EmailService
}

func NewAuthUsecase(
	userRepository repositories.UserRepository,
	jwtService services.JWTService,
	emailService services.EmailService,
) AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
		jwtService:     jwtService,
		emailService:   emailService,
	}
}

func (u *authUsecase) Register(input dto.RegisterUserInput) (types.JWTTokenResponse, types.JWTSession, error) {
	dbUser, err := u.userRepository.GetByEmail(input.Email)
	var tokens types.JWTTokenResponse
	var session types.JWTSession

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return tokens, session, err
		}
	}

	if dbUser.ID != "" {
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

	go func() {
		err := u.emailService.SendVerificationEmail(newUser.Email)
		if err != nil {
			log.Println(err.Error())
		}
	}()

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
		if errors.Is(err, pgx.ErrNoRows) {
			return tokens, session, errs.ErrUserDoesNotExist
		} else {
			return tokens, session, err
		}
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

func (u *authUsecase) VerifyEmail(email string) (types.JWTTokenResponse, types.JWTSession, error) {
	var tokens types.JWTTokenResponse
	var session types.JWTSession

	newUser := &entities.User{
		EmailVerified: true,
	}

	err := u.userRepository.UpdateByEmail(email, newUser)

	if err != nil {
		return tokens, session, errs.ErrFailedToVerifyEmail
	}

	tokens, session, err = u.GenerateSession(newUser)

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
