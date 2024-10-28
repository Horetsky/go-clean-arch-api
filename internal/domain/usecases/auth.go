package usecases

import (
	"errors"
	"log"
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"
	errs "seeker/internal/domain/errors"
	"seeker/internal/domain/repositories"
	"seeker/internal/domain/services"

	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(input dto.RegisterUserInput) (dto.JWTTokenResponse, dto.JWTSession, error)
	Login(input dto.LoginUserInput) (dto.JWTTokenResponse, dto.JWTSession, error)
	GenerateSession(user *entities.User) (dto.JWTTokenResponse, dto.JWTSession, error)
	VerifyEmail(email string) (dto.JWTTokenResponse, dto.JWTSession, error)
	DeleteAccount(email string) error
}

type authUsecase struct {
	userRepository      repositories.UserRepository
	talentRepository    repositories.TalentRepository
	recruiterRepository repositories.RecruiterRepository
	jwtService          services.JWTService
	emailService        services.EmailService
}

func NewAuthUsecase(
	userRepository repositories.UserRepository,
	talentRepository repositories.TalentRepository,
	recruiterRepository repositories.RecruiterRepository,
	jwtService services.JWTService,
	emailService services.EmailService,
) AuthUsecase {
	return &authUsecase{
		userRepository:      userRepository,
		talentRepository:    talentRepository,
		recruiterRepository: recruiterRepository,
		jwtService:          jwtService,
		emailService:        emailService,
	}
}

func (u *authUsecase) Register(input dto.RegisterUserInput) (dto.JWTTokenResponse, dto.JWTSession, error) {
	dbUser, err := u.userRepository.FindByEmail(input.Email)
	var tokens dto.JWTTokenResponse
	var session dto.JWTSession

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
		Type:  input.Type,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return tokens, session, errs.ErrFailedToCreateUser
	}

	newUser.Password = string(hashedPassword)

	err = u.userRepository.Create(newUser)

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

func (u *authUsecase) Login(input dto.LoginUserInput) (dto.JWTTokenResponse, dto.JWTSession, error) {
	dbUser, err := u.userRepository.FindByEmail(input.Email)
	var tokens dto.JWTTokenResponse
	var session dto.JWTSession

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

	if dbUser.Type == entities.TalentType {
		talent, err := u.talentRepository.FindByUserID(dbUser.ID)
		if err == nil {
			dbUser.Talent = &talent
		}
	}

	if dbUser.Type == entities.RecruiterType {
		recruiter, err := u.recruiterRepository.FindByUserID(dbUser.ID)
		if err == nil {
			dbUser.Recruiter = &recruiter
		}
	}

	tokens, session, err = u.GenerateSession(&dbUser)

	if err != nil {
		return tokens, session, err
	}

	return tokens, session, nil
}

func (u *authUsecase) VerifyEmail(email string) (dto.JWTTokenResponse, dto.JWTSession, error) {
	var tokens dto.JWTTokenResponse
	var session dto.JWTSession

	newUser := entities.User{
		EmailVerified: true,
	}

	err := u.userRepository.UpdateByEmail(email, &newUser)

	if err != nil {
		return tokens, session, errs.ErrFailedToVerifyEmail
	}

	if newUser.Type == entities.TalentType {
		talent, err := u.talentRepository.FindByUserID(newUser.ID)
		if err == nil {
			newUser.Talent = &talent
		}
	}

	if newUser.Type == entities.RecruiterType {
		recruiter, err := u.recruiterRepository.FindByUserID(newUser.ID)
		if err == nil {
			newUser.Recruiter = &recruiter
		}
	}

	tokens, session, err = u.GenerateSession(&newUser)

	if err != nil {
		return tokens, session, err
	}

	return tokens, session, nil
}

func (u *authUsecase) DeleteAccount(email string) error {
	err := u.userRepository.DeleteByEmail(email)

	if err != nil {
		return errs.ErrFailedToDeleteAccount
	}

	return nil
}

func (u *authUsecase) GenerateSession(user *entities.User) (dto.JWTTokenResponse, dto.JWTSession, error) {
	var tokens dto.JWTTokenResponse
	var session dto.JWTSession

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
