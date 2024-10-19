package dto

type UpdateUserInput struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Picture  string `json:"picture,omitempty"`
}
