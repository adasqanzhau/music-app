package routes

import (
	handlers "music-app/internal/delivery"
	"music-app/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(songService *services.SongService) *gin.Engine {
	router := gin.Default()

	songHandler := handlers.NewSongHandler(songService)

	songs := router.Group("/songs")
	{
		songs.POST("/", songHandler.CreateSong)
		songs.GET("/", songHandler.GetAllSongs)
		songs.GET("/:id", songHandler.GetSongByID)
		songs.PUT("/:id", songHandler.UpdateSong)
		songs.DELETE("/:id", songHandler.DeleteSong)
	}

	return router
}
