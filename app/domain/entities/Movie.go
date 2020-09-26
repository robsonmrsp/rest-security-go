package entities

import "time"

// PagerMovie ...
type PagerMovie struct {
}

// Movie ...
type Movie struct {
	ID          int       `db:"id" `
	Title       string    `db:"name"`
	ReleaseDate time.Time `db:"release_date"`
	Genre       Genre
}
