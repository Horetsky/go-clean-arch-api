package entities

type Talent struct {
	ID      string         `json:"id,omitempty"`
	UserID  string         `json:"userId,omitempty"`
	Profile *TalentProfile `json:"profile,omitempty"`
}

type TalentProfile struct {
	ID          string `json:"id,omitempty"`
	TalentId    string `json:"talentId,omitempty"`
	Category    string `json:"category,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Phone       string `json:"phone,omitempty"`
	LinkedInUrl string `json:"linkedInUrl,omitempty"`
	ResumeUrl   string `json:"resumeUrl,omitempty"`
	Photo       string `json:"photo,omitempty"`
}
