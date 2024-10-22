package entities

type Job struct {
	ID           string `json:"id,omitempty"`
	RecruiterID  string `json:"recruiterId,omitempty"`
	Category     string `json:"category,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Requirements string `json:"requirements,omitempty"`
}
type JobWithRecruiter struct {
	Job
	Recruiter Recruiter `json:"recruiter"`
}

type JobApplication struct {
	TalentID string `json:"talentId,omitempty"`
	JobID    string `json:"jobId,omitempty"`
}
