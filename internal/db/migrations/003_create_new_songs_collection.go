package migrations

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func Up003(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection("new_songs")

	song := map[string]interface{}{
		"title":  "New Song Title",
		"author": "New Author",
		"length": 180,
		"cover":  "some_cover_url",
	}

	_, err := collection.InsertOne(ctx, song)
	if err != nil {
		log.Println("Error inserting song into new_songs:", err)
		return err
	}

	log.Println("Migration 003_seed_song applied: inserted a sample song")
	return nil
}
