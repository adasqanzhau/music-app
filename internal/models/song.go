package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title         string             `bson:"title" json:"title"`
	Author        string             `bson:"author" json:"author"`
	Length        int                `bson:"length,omitempty" json:"length,omitempty"`
	Cover         string             `bson:"cover,omitempty" json:"cover,omitempty"`
	OwnerUsername string             `bson:"ownerUsername" json:"ownerUsername"`
}
