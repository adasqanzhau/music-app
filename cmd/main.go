package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"music-app/internal/repository"
	"music-app/internal/routes"
	"music-app/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
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

	songRepo := repository.NewSongRepository(database)

	songService := services.NewSongService(songRepo)

	router := gin.Default()
	routes.SetupRouter(songService).Run(":5000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Server is running on port:", port)
	router.Run(":" + port)
}
