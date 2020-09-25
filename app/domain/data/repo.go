package data

import (
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

type GenreRepository interface {
	SaveGenre(genre *entities.Genre) (*entities.Genre, error)
	GetGenre(ID int) (*entities.Genre, error)
}


type MovieRepository interface {
	SaveMovie(movie *entities.Movie) (*entities.Movie, error)
	GetMovie(ID int) (*entities.Movie, error)
	GetPageMovie(page int, pageSize int, order string, orderBy string, filter parser.Parameters) (*[]entities.Movie, error)
}
