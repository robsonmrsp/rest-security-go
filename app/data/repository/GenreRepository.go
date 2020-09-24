package repository

import (
	"fmt"
	"math/rand"

	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
)

type GenreRepository struct {
	dataBaseHelper *DataBaseHelper
}

func NewGenreRepository(dbh *DataBaseHelper) *GenreRepository {
	repo := new(GenreRepository)
	repo.dataBaseHelper = dbh
	return repo
}

// SaveGenre ..
func (repo *GenreRepository) SaveGenre(genre *entities.Genre) (*entities.Genre, error) {
	newID := rand.Intn(10000)
	stmt, err := repo.dataBaseHelper.tx.Prepare(`INSERT INTO genre(id, title, release_date) VALUES( $1, $2 )`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	if _, err := stmt.Exec(newID, genre.Name); err != nil {
		return nil, err
	}
	genre.ID = newID

	return genre, nil
}

// GetGenre ...
func (repo *GenreRepository) GetGenre(ID int) (*entities.Genre, error) {
	mo := entities.Genre{}

	row := repo.dataBaseHelper.db.QueryRow("SELECT id, title, release_date age FROM genre where id = $1", ID)

	err := row.Scan(&mo.ID, &mo.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &mo, nil
}

// AllGenres ..
func AllGenres() ([]*entities.Genre, error) {
	dataBaseHelper := GetDb()
	rows, err := dataBaseHelper.db.Query("SELECT name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]*entities.Genre, 0)
	for rows.Next() {
		genre := new(entities.Genre)
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}
		people = append(people, genre)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return people, nil
}
