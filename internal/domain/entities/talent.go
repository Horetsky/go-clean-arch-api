package entities

type Talent struct {
	ID     string `json:"id,omitempty"`
	UserID string `json:"userId,omitempty"`
}

type TalentProfile struct {
	ID          string `json:"id,omitempty"`
	UserID      string `json:"userId,omitempty"`
	TalentId    string `json:"talentId,omitempty"`
	CategoryId  string `json:"categoryId,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Phone       string `json:"phone,omitempty"`
	LinkedInUrl string `json:"linkedInUrl,omitempty"`
	ResumeUrl   string `json:"resumeUrl,omitempty"`
	Photo       string `json:"photo,omitempty"`
}
