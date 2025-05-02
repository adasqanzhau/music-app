package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"music-app/internal/db"
	"music-app/internal/repository"
	"music-app/internal/routes"
	"music-app/internal/services"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB is not responding:", err)
	}
	log.Println("Connected to MongoDB")

	database := client.Database("music-db")

	if err := db.RunMigrations(database); err != nil {
		log.Fatal("Migrations failed:", err)
	}

	songRepo := repository.NewSongRepository(database)
	songService := services.NewSongService(songRepo)
	userRepo := repository.NewUserRepository(database)
	authService := services.NewAuthService(userRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router := routes.SetupRouter(songService, authService)
	log.Println("Server is running on port:", port)
	router.Run(":" + port)
}
