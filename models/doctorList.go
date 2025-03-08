package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DoctorList struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email`
	Designation string             `bson:"designation" json:"designation"`
}
