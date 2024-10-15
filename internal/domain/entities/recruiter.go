package entities

type Recruiter struct {
	ID     string `json:"id,omitempty"`
	UserID string `json:"userId,omitempty"`
}

type RecruiterProfile struct {
	ID                string `json:"id,omitempty"`
	RecruiterID       string `json:"recruiterId,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Phone             string `json:"phone,omitempty"`
	LinkedInUrl       string `json:"linkedInUrl,omitempty"`
	CompanyName       string `json:"companyName,omitempty"`
	CompanyWebsiteUrl string `json:"companyWebsiteUrl,omitempty"`
}
