package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie        string             `json:"movie,omitempty" bson:"name,omitempty"`
	watched bool `json:"watched,omitempty" bson:"watched,omitempty"`
}