package service

import (
    "context"
    "github.com/track-service/internal/model"
    "github.com/track-service/internal/repository"
)

type TrackService struct {
    repo *repository.TrackRepository
}

func NewTrackService(repo *repository.TrackRepository) *TrackService {
    return &TrackService{repo: repo}
}

func (s *TrackService) GetTrack(ctx context.Context, id string) (*model.Track, error) {
    return s.repo.GetTrackByID(ctx, id)
}

func (s *TrackService) CreateTrack(ctx context.Context, track *model.Track) error {
    return s.repo.CreateTrack(ctx, track)
}