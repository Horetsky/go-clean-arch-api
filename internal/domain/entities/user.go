package entities

type User struct {
	ID            string  `json:"id,omitempty"`
	Email         string  `json:"email,omitempty"`
	Password      string  `json:"password,omitempty"`
	Picture       *string `json:"picture,omitempty"`
	EmailVerified bool    `json:"emailVerified"`
	Type          string  `json:"type,omitempty"`

	Talent    *Talent    `json:"talent,omitempty"`
	Recruiter *Recruiter `json:"recruiter,omitempty"`
}

const (
	TalentType    = "TALENT"
	RecruiterType = "RECRUITER"
)
