package routes

import (
	handlers "music-app/internal/delivery"
	"music-app/internal/middleware"
	"music-app/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(songService *services.SongService, authService *services.AuthService) *gin.Engine {
	router := gin.Default()

	authHandler := handlers.NewAuthHandler(authService)
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	songHandler := handlers.NewSongHandler(songService)

	songs := router.Group("/songs")
	songs.Use(middleware.AuthMiddleware())
	{
		songs.POST("/", songHandler.CreateSong)
		songs.GET("/:id", songHandler.GetSongByID)
		songs.PUT("/:id", songHandler.UpdateSong)
		songs.DELETE("/:id", songHandler.DeleteSong)
		songs.GET("/my-songs", songHandler.GetMySongs)
		songs.GET("/all-songs", songHandler.GetAllSongsForAdmin)
	}

	return router
}
