package services

import (
	"errors"
	"log"

	"github.com/robsonmrsp/rest-security-go/app/domain/data"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

type movieService struct {
	repo data.MovieRepository
}

func NewMovieService(repository data.MovieRepository) MovieService {
	return &movieService{repository}
}

func (serv *movieService) SaveMovie(movie *entities.Movie) (*entities.Movie, error) {
	movie, err := serv.repo.SaveMovie(movie)
	if err != nil {
		log.Print("Error saving movie", err)
		return nil, err
	}

	return movie, nil
}

func (serv *movieService) GetMovie(id int) (*entities.Movie, error) {
	m, err := serv.repo.GetMovie(id)

	if err != nil {
		return &entities.Movie{}, errors.New("Movie not find")
	}
	return m, nil
}

// Possible problem with zeros

func (serv *movieService) GetPageMovie(p *parser.PageParameters) (*[]entities.Movie, error) {
	m, err := serv.repo.GetPageMovie(p.Page, p.PageSize, p.Order, p.OrderBy, *p.FilterParameters)

	if err != nil {
		return &[]entities.Movie{}, errors.New("Movie not find")
	}
	return m, nil
}
