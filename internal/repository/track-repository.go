// internal/repository/track_repository.go
package repository

import (
    "context"
    "github.com/jackc/pgx/v4/pgxpool"
    "track-service/internal/model"
)

type TrackRepository struct {
    pool *pgxpool.Pool
}

func NewTrackRepository(pool *pgxpool.Pool) *TrackRepository {
    return &TrackRepository{pool: pool}
}

func (r *TrackRepository) GetTrackByID(ctx context.Context, id string) (*model.Track, error) {
    var track model.Track
    err := r.pool.QueryRow(ctx, "SELECT id, title, album_id, artist_id, duration_sec, genre, release_date, url FROM tracks WHERE id = $1", id).Scan(
        &track.ID,
        &track.Title,
        &track.AlbumID,
        &track.ArtistID,
        &track.DurationSec,
        &track.Genre,
        &track.ReleaseDate,
        &track.URL,
    )
    if err != nil {
        return nil, err
    }
    return &track, nil
}

func (r *TrackRepository) CreateTrack(ctx context.Context, track *model.Track) error {
    _, err := r.pool.Exec(ctx, `
        INSERT INTO tracks (id, title, album_id, artist_id, duration_sec, genre, release_date, url)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
        track.ID,
        track.Title,
        track.AlbumID,
        track.ArtistID,
        track.DurationSec,
        track.Genre,
        track.ReleaseDate,
        track.URL,
    )
    return err
}

// Добавь UpdateTrack, DeleteTrack и другие методы самостоятельно