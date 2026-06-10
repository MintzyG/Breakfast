package notes

import (
	"context"
	"pancake/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	col *mongo.Collection
}

func NewRepository(col *mongo.Collection) *Repository {
	return &Repository{col: col}
}

func (r *Repository) FindAll(ctx context.Context, filter bson.M) ([]models.Note, error) {
	opts := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}})
	cursor, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notes []models.Note
	if err := cursor.All(ctx, &notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *Repository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Note, error) {
	var note models.Note
	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *Repository) Create(ctx context.Context, note *models.Note) error {
	note.ID = primitive.NewObjectID()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	_, err := r.col.InsertOne(ctx, note)
	return err
}

func (r *Repository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	update["updated_at"] = time.Now()
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
