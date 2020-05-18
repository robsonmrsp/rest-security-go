package services

import (
	"errors"
	"log"

	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
)

type MovieService interface {
	SaveMovie(movie *entities.Movie) (*entities.Movie, error)
	GetMovie(id int) (*entities.Movie, error)
}

type Repository interface {
	SaveMovie(movie *entities.Movie) (*entities.Movie, error)
	GetMovie(id int) (*entities.Movie, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) MovieService {
	return &service{repository}
}

func (serv *service) SaveMovie(movie *entities.Movie) (*entities.Movie, error) {
	movie, err := serv.repo.SaveMovie(movie)
	if err != nil {
		log.Print("Error saving movie", err)
		return nil, err
	}

	return movie, nil
}

func (serv *service) GetMovie(id int) (*entities.Movie, error) {
	m, err := serv.repo.GetMovie(id)

	if err != nil {
		return &entities.Movie{}, errors.New("Movie not find")
	}
	return m, nil
}
