package migrations

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Up001(ctx context.Context, db *mongo.Database) error {
	songs := db.Collection("new_songs")

	_, err := songs.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"title": 1},
	})
	if err != nil {
		log.Println("Error creating index:", err)
		return err
	}

	log.Println("Migration 001_init applied: created index on 'title'")
	return nil
}
