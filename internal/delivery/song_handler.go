package handlers

import (
	"context"
	"music-app/internal/models"
	"music-app/internal/services"
	"net/http"

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

	username := c.GetString("username")
	song.OwnerUsername = username

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

	username := c.GetString("username")
	role := c.GetString("role")

	if role != "admin" && song.OwnerUsername != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *SongHandler) GetMySongs(c *gin.Context) {
	username := c.GetString("username")

	songs, err := h.service.GetSongsByOwner(context.Background(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *SongHandler) UpdateSong(c *gin.Context) {
	id := c.Param("id")

	song, err := h.service.GetSongByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	username := c.GetString("username")
	role := c.GetString("role")

	if role != "admin" && song.OwnerUsername != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var updated models.Song
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated.ID = song.ID
	updated.OwnerUsername = song.OwnerUsername

	err = h.service.UpdateSong(context.Background(), id, &updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *SongHandler) DeleteSong(c *gin.Context) {
	id := c.Param("id")

	song, err := h.service.GetSongByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	username := c.GetString("username")
	role := c.GetString("role")

	if role != "admin" && song.OwnerUsername != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	err = h.service.DeleteSong(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
}

func (h *SongHandler) GetAllSongsForAdmin(c *gin.Context) {
	role := c.GetString("role")

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	songs, err := h.service.GetAllSongsForAdmin(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}
