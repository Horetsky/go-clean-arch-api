package dto

type CreateRecruiterProfileInput struct {
	UserID            string `json:"userId,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Phone             string `json:"phone,omitempty"`
	LinkedInUrl       string `json:"linkedInUrl,omitempty"`
	CompanyName       string `json:"companyName,omitempty"`
	CompanyWebsiteUrl string `json:"companyWebsiteUrl,omitempty"`
}

type UpdateRecruiterProfile struct {
	*CreateRecruiterProfileInput
}
