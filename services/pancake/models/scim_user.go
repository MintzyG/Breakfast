package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SCIMUserSchema = "urn:ietf:params:scim:schemas:core:2.0:User"

type SCIMUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Schemas  []string           `bson:"schemas"       json:"schemas"`
	UserName string             `bson:"userName"      json:"userName"`
	Name     SCIMName           `bson:"name"          json:"name"`
	Emails   []SCIMEmail        `bson:"emails"        json:"emails"`
	Active   bool               `bson:"active"        json:"active"`
	Meta     SCIMUserMeta       `bson:"meta"          json:"meta"`
}

type SCIMName struct {
	GivenName  string `bson:"givenName"  json:"givenName"`
	FamilyName string `bson:"familyName" json:"familyName"`
}

type SCIMEmail struct {
	Value   string `bson:"value"   json:"value"`
	Primary bool   `bson:"primary" json:"primary"`
}

type SCIMUserMeta struct {
	ResourceType string    `bson:"resourceType" json:"resourceType"`
	Created      time.Time `bson:"created"      json:"created"`
	LastModified time.Time `bson:"lastModified" json:"lastModified"`
}

type SCIMListResponse struct {
	Schemas      []string   `json:"schemas"`
	TotalResults int        `json:"totalResults"`
	StartIndex   int        `json:"startIndex"`
	ItemsPerPage int        `json:"itemsPerPage"`
	Resources    []SCIMUser `json:"Resources"`
}

type SCIMPatchOp struct {
	Schemas    []string        `json:"schemas"`
	Operations []SCIMOperation `json:"Operations"`
}

type SCIMOperation struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value any    `json:"value"`
}
