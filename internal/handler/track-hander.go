// internal/handler/track_handler.go
package handler

import (
    "context"
    "github.com/vadkvadrad/track-service/internal/model"
    "github.com/vadkvadrad/track-service/internal/service"
    pb "github.com/vadkvadrad/track-service/proto"
)

type TrackHandler struct {
    pb.UnimplementedTrackServiceServer
    service *service.TrackService
}

func NewTrackHandler(service *service.TrackService) *TrackHandler {
    return &TrackHandler{service: service}
}

func (h *TrackHandler) GetTrack(ctx context.Context, req *pb.TrackIDRequest) (*pb.TrackResponse, error) {
    track, err := h.service.GetTrack(ctx, req.Id)
    if err != nil {
        return nil, err
    }
    return &pb.TrackResponse{
        Id:           track.ID,
        Title:        track.Title,
        AlbumId:      track.AlbumID,
        ArtistId:     track.ArtistID,
        DurationSec:  int32(track.DurationSec),
        Genre:        track.Genre,
        ReleaseDate:  track.ReleaseDate,
        Url:          track.URL,
    }, nil
}

func (h *TrackHandler) CreateTrack(ctx context.Context, req *pb.CreateTrackRequest) (*pb.Empty, error) {
    track := &model.Track{
        ID:          req.Id,
        Title:       req.Title,
        AlbumID:     req.AlbumId,
        ArtistID:    req.ArtistId,
        DurationSec: int(req.DurationSec),
        Genre:       req.Genre,
        ReleaseDate: req.ReleaseDate,
        URL:         req.Url,
    }
    return &pb.Empty{}, h.service.CreateTrack(ctx, track)
}