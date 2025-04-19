package migrations

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Up002(ctx context.Context, db *mongo.Database) error {
	users := db.Collection("users")

	_, err := users.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: nil,
	})
	if err != nil {
		log.Println("Error creating index on users:", err)
		return err
	}

	log.Println("Migration 002_create_users_collection applied: created index on 'username'")
	return nil
}
