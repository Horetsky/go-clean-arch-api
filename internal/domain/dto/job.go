package dto

type PostJobDTO struct {
	RecruiterID  string `json:"recruiterId,omitempty"`
	Category     string `json:"category,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Requirements string `json:"requirements,omitempty"`
}

type UpdateJobDTO struct {
	ID           string `json:"id,omitempty"`
	RecruiterID  string `json:"recruiterId,omitempty"`
	Category     string `json:"category,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Requirements string `json:"requirements,omitempty"`
}
