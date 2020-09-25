package services

import (
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

// GenreService ...
type GenreService interface {
	SaveGenre(genre *entities.Genre) (*entities.Genre, error)
	GetGenre(id int) (*entities.Genre, error)
}

// MovieService ...
type MovieService interface {
	SaveMovie(movie *entities.Movie) (*entities.Movie, error)
	GetMovie(id int) (*entities.Movie, error)
	GetPageMovie(*parser.PageParameters) (*[]entities.Movie, error)
}
