package handlers

import (
	"context"
	"net/http"

	"music-app/internal/models"
	"music-app/internal/services"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
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
