package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seeker/internal/domain/dto"
)

func GetContextValue(r *http.Request, key string) any {
	ctx := r.Context()
	return ctx.Value(key)
}

func GetSession(r *http.Request) (*dto.JWTSession, error) {
	user, ok := GetContextValue(r, dto.CtxSessionKey).(*dto.JWTUser)

	if !ok || user == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	session := dto.JWTSession{
		User: *user,
	}

	return &session, nil
}

func ReadBody(r *http.Request, p any) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		return err
	}

	return nil
}
