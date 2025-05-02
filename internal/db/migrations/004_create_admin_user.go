package migrations

import (
	"context"
	"log"
	"os"

	"music-app/internal/models"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Up004(ctx context.Context, db *mongo.Database) error {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	if adminUsername == "" {
		adminUsername = "admin"
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "admin123"
	}

	users := db.Collection("users")

	var adminUser models.User
	err := users.FindOne(ctx, bson.M{"username": adminUsername}).Decode(&adminUser)
	if err == nil {
		log.Println("Admin user already exists, skipping creation")
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = users.InsertOne(ctx, &models.User{
		Username: adminUsername,
		Password: string(hashedPassword),
		Role:     "admin",
	})
	if err != nil {
		return err
	}

	log.Printf("Admin user '%s' created successfully", adminUsername)
	return nil
}
