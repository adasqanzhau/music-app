package handlers

import (
	"context"
	"errors"
	"net/http"

	"music-app/internal/models"
	"music-app/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type SongHandler struct {
	service *services.SongService
}

func NewSongHandler(service *services.SongService) *SongHandler {
	return &SongHandler{service: service}
}

func (h *SongHandler) CreateSong(c *gin.Context) {
	var song models.Song

	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if song.Title == "" || song.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Author are required"})
		return
	}

	createdSong, err := h.service.CreateSong(context.Background(), &song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSong)
}

func (h *SongHandler) GetSongByID(c *gin.Context) {
	id := c.Param("id")

	song, err := h.service.GetSongByID(context.Background(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *SongHandler) GetAllSongs(c *gin.Context) {
	songs, err := h.service.GetAllSongs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *SongHandler) UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var updatedSong models.Song

	if err := c.ShouldBindJSON(&updatedSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateSong(context.Background(), id, &updatedSong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully"})
}

func (h *SongHandler) DeleteSong(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteSong(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
}
