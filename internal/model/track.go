package model

type Track struct {
    ID          string `json:"id" db:"id"`
    Title       string `json:"title" db:"title"`
    AlbumID     string `json:"album_id" db:"album_id"`
    ArtistID    string `json:"artist_id" db:"artist_id"`
    DurationSec int    `json:"duration_sec" db:"duration_sec"`
    Genre       string `json:"genre" db:"genre"`
    ReleaseDate string `json:"release_date" db:"release_date"`
    URL         string `json:"url" db:"url"`
}