package entities

import "time"

// PagerMovie ...
type PagerMovie struct {
	Items        *[]Movie `json:"items"`
	ActualPage   int      `json:"actualPage"`
	TotalRecords int32    `json:"totalRecords"`

	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
}

// Movie ...
type Movie struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"name" json:"title"`
	ReleaseDate time.Time `db:"release_date" json:"releaseDate"`
	Genre       *Genre    `db:"genre_id" json:"genre"`
}
