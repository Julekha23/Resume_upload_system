package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReviewerStatus string

const (
	ReviewerActive   ReviewerStatus = "active"
	ReviewerInactive ReviewerStatus = "inactive"
)

type Reviewer struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ReviewerID string `bson:"reviewer_id" json:"reviewer_id"`
	Name        string             `bson:"name"          json:"name"`
	Email       string             `bson:"email"         json:"email"`
	Status      ReviewerStatus     `bson:"status"        json:"status"`
}