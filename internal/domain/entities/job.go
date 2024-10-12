package entities

type Job struct {
	ID           string `json:"is,omitempty"`
	RecruiterID  string `json:"recruiterId,omitempty"`
	CategoryID   string `json:"categoryId,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Requirements string `json:"requirements,omitempty"`
}
