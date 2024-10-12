package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seeker/internal/types"
)

func GetContextValue(r *http.Request, key string) any {
	ctx := r.Context()
	return ctx.Value(key)
}

func GetSession(r *http.Request) (*types.JWTUser, error) {
	session, ok := GetContextValue(r, types.CtxSessionKey).(*types.JWTUser)

	if !ok || session == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	return session, nil
}

func ReadBody(r *http.Request, p any) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		return err
	}

	return nil
}
