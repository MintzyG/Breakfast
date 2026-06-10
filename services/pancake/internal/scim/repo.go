package scim

import (
	"context"
	"pancake/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	col *mongo.Collection
}

func NewRepository(col *mongo.Collection) *Repository {
	return &Repository{col: col}
}

func (r *Repository) FindAll(ctx context.Context, filter bson.M) ([]models.SCIMUser, error) {
	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.SCIMUser
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	if users == nil {
		users = []models.SCIMUser{}
	}
	return users, nil
}

func (r *Repository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.SCIMUser, error) {
	var user models.SCIMUser
	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Create(ctx context.Context, user *models.SCIMUser) error {
	user.ID = primitive.NewObjectID()
	now := time.Now()
	user.Meta = models.SCIMUserMeta{
		ResourceType: "User",
		Created:      now,
		LastModified: now,
	}
	if len(user.Schemas) == 0 {
		user.Schemas = []string{models.SCIMUserSchema}
	}
	_, err := r.col.InsertOne(ctx, user)
	return err
}

func (r *Repository) Replace(ctx context.Context, id primitive.ObjectID, user *models.SCIMUser) error {
	user.ID = id
	user.Meta.LastModified = time.Now()
	if len(user.Schemas) == 0 {
		user.Schemas = []string{models.SCIMUserSchema}
	}
	_, err := r.col.ReplaceOne(ctx, bson.M{"_id": id}, user)
	return err
}

func (r *Repository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	update["meta.lastModified"] = time.Now()
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
