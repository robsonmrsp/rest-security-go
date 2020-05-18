package entities

import "time"

type PagerMovie struct {
	
}

type Movie struct {
	ID          int       `db:"id" `
	Title       string    `db:"name"`
	ReleaseDate time.Time `db:"release_date"`
	Genres      []Genre
}
