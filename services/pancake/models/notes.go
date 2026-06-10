package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"    json:"id"`
	Title     string             `bson:"title"            json:"title"`
	Content   []byte             `bson:"content"          json:"content"`
	Color     string             `bson:"color"            json:"color"`
	Emoji     string             `bson:"emoji"            json:"emoji"`
	Tags      []string           `bson:"tags"             json:"tags"`
	CreatedAt time.Time          `bson:"created_at"       json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at"       json:"updatedAt"`
}
