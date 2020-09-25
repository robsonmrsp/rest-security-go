package services

import (
	"errors"
	"log"

	"github.com/robsonmrsp/rest-security-go/app/domain/data"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
)

type genreService struct {
	repo data.GenreRepository
}

func NewGenreService(repository data.GenreRepository) GenreService {
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
