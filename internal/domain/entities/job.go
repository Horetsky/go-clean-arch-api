package entities

type Job struct {
	ID           string `json:"is,omitempty"`
	RecruiterID  string `json:"recruiterId,omitempty"`
	Category     string `json:"category,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Requirements string `json:"requirements,omitempty"`
}
