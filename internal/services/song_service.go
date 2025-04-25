package services

import (
	"context"
	"music-app/internal/models"
	"music-app/internal/repository"
	"time"
)

type SongService struct {
	repo *repository.SongRepository
}

func NewSongService(repo *repository.SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(ctx context.Context, song *models.Song) (*models.Song, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.CreateSong(ctx, song)
}

func (s *SongService) GetSongByID(ctx context.Context, id string) (*models.Song, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.GetSongByID(ctx, id)
}

func (s *SongService) GetAllSongs(ctx context.Context) ([]models.Song, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.GetAllSongs(ctx)
}

func (s *SongService) GetAllSongsForAdmin(ctx context.Context) ([]models.Song, error) {
	return s.repo.GetAllSongs(ctx)
}

func (s *SongService) UpdateSong(ctx context.Context, id string, updatedSong *models.Song) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.UpdateSong(ctx, id, updatedSong)
}

func (s *SongService) DeleteSong(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.DeleteSong(ctx, id)
}

func (s *SongService) GetSongsByOwner(ctx context.Context, ownerUsername string) ([]*models.Song, error) {
	return s.repo.GetSongsByOwner(ctx, ownerUsername)
}
