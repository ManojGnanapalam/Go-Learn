package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Name    string
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Watched bool
}
