package db

import (
	"context"
	"log"
	"time"

	"music-app/internal/db/migrations"

	"go.mongodb.org/mongo-driver/mongo"
)

type Migration struct {
	ID string
	Up func(context.Context, *mongo.Database) error
}

var allMigrations = []Migration{
	{"001_init", migrations.Up001},
	{"002_create_users_collection", migrations.Up002},
	{"003_seed_song", migrations.Up003},
	{"004_create_admin_user", migrations.Up004},
}

func RunMigrations(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	applied := make(map[string]bool)
	cursor, err := db.Collection("migrations").Find(ctx, map[string]interface{}{})
	if err == nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var m struct{ ID string }
			_ = cursor.Decode(&m)
			applied[m.ID] = true
		}
	}

	for _, m := range allMigrations {
		if !applied[m.ID] {
			log.Printf("Running migration: %s...\n", m.ID)
			if err := m.Up(ctx, db); err != nil {
				return err
			}
			_, _ = db.Collection("migrations").InsertOne(ctx, map[string]interface{}{
				"id":        m.ID,
				"appliedAt": time.Now(),
			})
		}
	}

	log.Println("All migrations applied.")
	return nil
}
