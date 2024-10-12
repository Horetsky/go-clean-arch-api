package dto

type RegisterUserInput struct {
	Email    string
	Password string
	Type     string // talent / recruiter
}

type LoginUserInput struct {
	Email    string
	Password string
}
type UpdateUserInput struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Picture  string `json:"picture,omitempty"`
}
