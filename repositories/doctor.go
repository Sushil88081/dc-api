package repository

import (
	"context"
	"doctor-on-demand/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDoctorRepository interface {
	CreateDoctor(user models.DoctorList) error
	GetDoctorByID(id string) (*models.DoctorList, error)
	UpdateDoctor(id string, userData map[string]interface{}) error
	DeleteDoctor(id string) error
}

type DoctorRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) IDoctorRepository {
	return &DoctorRepository{
		collection: db.Collection("Doctors"),
	}
}

func (r *DoctorRepository) CreateDoctor(user models.DoctorList) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}
func (r *DoctorRepository) GetDoctorByID(id string) (*models.DoctorList, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid doctor ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.DoctorList
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("doctor not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// 🔹 Update user (partial update)
func (r *DoctorRepository) UpdateDoctor(id string, userData map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid doctor ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateData := bson.M{"$set": userData}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, updateData)
	if err != nil {
		log.Println("UpdateDoctor Error:", err)
		return err
	}

	return nil
}

// 🔹 Delete user
func (r *DoctorRepository) DeleteDoctor(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Println("Delete Doctor Error:", err)
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("Doctor not found")
	}

	return nil
}
