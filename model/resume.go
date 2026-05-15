package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResumeStatus string
const (
	StatusPending   ResumeStatus = "pending"
	StatusReviewing ResumeStatus = "reviewing"
	StatusDone      ResumeStatus = "done"
	StatusRejected  ResumeStatus = "rejected"
)
type Resume struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"        json:"-"`
	ResumeID         string             `bson:"resume_id"            json:"resume_id"`
	CandidateName    string             `bson:"candidate_name"       json:"candidate_name"`
	CandidateEmail   string             `bson:"candidate_email"      json:"candidate_email"`
	Phone            string             `bson:"phone,omitempty"      json:"phone,omitempty"`
	OriginalFileName string             `bson:"original_file_name"   json:"original_file_name"`
	FilePath         string             `bson:"file_path"            json:"file_path"`
	FileSizeBytes    int64              `bson:"file_size_bytes"      json:"file_size_bytes"`
	Status           ResumeStatus       `bson:"status,omitempty"               json:"status"`
	ReviewerID       string             `bson:"reviewer_id,omitempty"          json:"reviewer_id"`
}
