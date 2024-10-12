package entities

type User struct {
	ID       string  `json:"id,omitempty"`
	Email    string  `json:"email,omitempty"`
	Password string  `json:"password,omitempty"`
	Picture  *string `json:"picture,omitempty"`
}
