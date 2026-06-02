package model
type UploadResponse struct {
	ResumeID string `json:"resume_id"` 
	Status   string `json:"status"`    
	Message  string `json:"message"`   
}
type StatusResponse struct {
	ResumeID   string `json:"resume_id"`
	Status     string `json:"status"`
	ReviewerID string `json:"reviewer_id,omitempty"`
	Message    string `json:"message"`
}