package types

import "github.com/golang-jwt/jwt/v5"

type JWTTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTSession struct {
	User JWTUser `json:"user,omitempty"`
	jwt.RegisteredClaims
}

type JWTUser struct {
	ID            string  `json:"id,omitempty"`
	Email         string  `json:"email,omitempty"`
	Picture       *string `json:"picture,omitempty"`
	EmailVerified bool    `json:"emailVerified"`
}

const (
	AccessTokenCookieKey  = "access_token"
	RefreshTokenCookieKey = "refresh_token"
	CtxSessionKey         = "session"
)
