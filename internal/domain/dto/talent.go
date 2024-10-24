package dto

type CreateTalentProfileInput struct {
	UserID      string `json:"userId,omitempty"`
	Category    string `json:"category,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Phone       string `json:"phone,omitempty"`
	LinkedInUrl string `json:"linkedInUrl,omitempty"`
	ResumeUrl   string `json:"resumeUrl,omitempty"`
	Photo       string `json:"photo,omitempty"`
}

type UpdateTalentProfile struct {
	*CreateTalentProfileInput
}

type ListTalentDTO struct {
	Category string `json:"category,omitempty"`
}
