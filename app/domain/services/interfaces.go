package services

import (
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

// GenreService ...
type GenreService interface {
	GetGenre(id int) (*entities.Genre, error)
	SaveGenre(genre *entities.Genre) (*entities.Genre, error)
	UpdateGenre(genre *entities.Genre) (*entities.Genre, error)
	DeleteGenre(ID int32) (bool, error)
}

// MovieService ...
type MovieService interface {
	SaveMovie(movie *entities.Movie) (*entities.Movie, error)
	UpdateMovie(movie *entities.Movie) (*entities.Movie, error)
	GetMovie(id int) (*entities.Movie, error)
	DeleteMovie(id int) (bool, error)
	GetPageMovie(*parser.PageParameters) (*entities.PagerMovie, error)
}
