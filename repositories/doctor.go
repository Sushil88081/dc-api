package repository

import (
	"context"
	"doctor-on-demand/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDoctorRepository interface {
	CreateDoctor(ctx context.Context, doctor models.DoctorList) error
	GetDoctor(ctx context.Context, id string) (models.DoctorList, error)
	UpdateDoctor(ctx context.Context, id string, doctor models.DoctorList) error
	DeleteDoctor(ctx context.Context, id string) error // ✅ यह method add करो
}

type DoctorRepository struct {
	collection *mongo.Collection
}

func NewDoctorRepository(collection *mongo.Collection) *DoctorRepository {
	return &DoctorRepository{collection: collection}
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doctor models.DoctorList) error {
	_, err := r.collection.InsertOne(ctx, doctor)
	return err
}

func (r *DoctorRepository) GetDoctor(ctx context.Context, id string) (models.DoctorList, error) {
	var doctor models.DoctorList
	err := r.collection.FindOne(ctx, id).Decode(&doctor)
	return doctor, err
}

func (r *DoctorRepository) UpdateDoctor(ctx context.Context, id string, doctor models.DoctorList) error {
	_, err := r.collection.UpdateOne(ctx, id, doctor)
	return err
}

// ✅ अब `DeleteDoctor` को implement किया:
func (r *DoctorRepository) DeleteDoctor(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, id)
	return err
}
