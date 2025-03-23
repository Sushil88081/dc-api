package models

type DoctorList struct {
	ID          string `bson:"_id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Email       string `bson:"email" json:"email`
	Designation string `bson:"designation" json:"designation"`
}
