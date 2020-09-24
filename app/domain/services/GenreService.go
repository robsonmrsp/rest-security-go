package services

import (
	"errors"
	"log"

	"github.com/robsonmrsp/rest-security-go/app/data/repository"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
)

type GenreService interface {
	SaveGenre(genre *entities.Genre) (*entities.Genre, error)
	GetGenre(id int) (*entities.Genre, error)
}

type genreService struct {
	repo repository.GenreRepository
}

func NewGenreService(repository repository.GenreRepository) GenreService {
	return &genreService{repository}
}

func (serv *genreService) SaveGenre(genre *entities.Genre) (*entities.Genre, error) {
	genre, err := serv.repo.SaveGenre(genre)
	if err != nil {
		log.Print("Error saving genre", err)
		return nil, err
	}

	return genre, nil
}

func (serv *genreService) GetGenre(id int) (*entities.Genre, error) {
	m, err := serv.repo.GetGenre(id)

	if err != nil {
		return &entities.Genre{}, errors.New("Genre not find")
	}
	return m, nil
}
