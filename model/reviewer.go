package model

import "go.mongodb.org/mongo-driver/bson/primitive"
<<<<<<< HEAD
type ReviewerStatus string
const (
	ReviewerActive   ReviewerStatus = "active"   
	ReviewerInactive ReviewerStatus = "inactive" 
)

type reviewer struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Reviewer_ID string `bson:"reviewer_id"   json:"reviewer_id"`
	Name  string `bson:"name"  json:"name"`
	Email string `bson:"email" json:"email"`
	Status ReviewerStatus `bson:"status" json:"status"`
} 
=======

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
>>>>>>> a1f77ceab5fe9317d46f796ad157c7ae4e4b40e1
