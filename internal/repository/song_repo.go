package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"music-app/internal/models"
)

type SongRepository struct {
	collection *mongo.Collection
}

func NewSongRepository(db *mongo.Database) *SongRepository {
	return &SongRepository{
		collection: db.Collection("new_songs"),
	}
}

func (r *SongRepository) CreateSong(ctx context.Context, song *models.Song) (*models.Song, error) {
	_, err := r.collection.InsertOne(ctx, song)
	if err != nil {
		return nil, err
	}
	return song, nil
}

func (r *SongRepository) GetSongByID(ctx context.Context, id string) (*models.Song, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var song models.Song
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&song)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *SongRepository) GetAllSongs(ctx context.Context) ([]models.Song, error) {
	var songs []models.Song
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var song models.Song
		if err := cursor.Decode(&song); err != nil {
			log.Println("Error decoding song:", err)
			continue
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongRepository) UpdateSong(ctx context.Context, id string, updatedSong *models.Song) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": updatedSong}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *SongRepository) DeleteSong(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *SongRepository) GetSongsByOwner(ctx context.Context, ownerUsername string) ([]*models.Song, error) {
	filter := bson.M{"ownerUsername": ownerUsername}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var songs []*models.Song
	for cursor.Next(ctx) {
		var song models.Song
		if err := cursor.Decode(&song); err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	return songs, nil
}
