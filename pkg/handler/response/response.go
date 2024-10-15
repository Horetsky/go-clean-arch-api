package response

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	JSON(w http.ResponseWriter, data any, status int)
	Error(w http.ResponseWriter, err error, status int)
}

type errorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type jsonResponse struct {
	Data any `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := jsonResponse{
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}

func Error(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var response errorResponse

	if err == nil {
		response.Error.Message = "Bad request"
	} else {
		response.Error.Message = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}

func PrivateCookie(w http.ResponseWriter, key string, value string) {
	cookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   5 * 60,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
}
