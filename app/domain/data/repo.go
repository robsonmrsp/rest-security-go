package data

import (
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

type GenreRepository interface {
	SaveGenre(genre *entities.Genre) (*entities.Genre, error)
	GetGenre(ID int) (*entities.Genre, error)
}

// MovieRepository ...
type MovieRepository interface {
	GetMovie(ID int) (*entities.Movie, error)
	SaveMovie(movie *entities.Movie) (*entities.Movie, error)
	DeleteMovie(movie *entities.Movie) (*entities.Movie, error)
	UpdateMovie(movie *entities.Movie) (*entities.Movie, error)
	GetPageMovie(page int, pageSize int, order string, orderBy string, filter parser.Parameters) (*[]entities.Movie, error)
	GetTotalMovies(page int, pageSize int, order string, orderBy string, filter parser.Parameters) (int32, error)
}
