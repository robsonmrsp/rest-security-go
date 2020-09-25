package repository

import (
	"fmt"
	"math/rand"

	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
)

type GenreRepository interface {
	SaveGenre(genre *entities.Genre) (*entities.Genre, error)
	GetGenre(ID int) (*entities.Genre, error)
}

type genreRepository struct {
	dataBaseHelper *DataBaseHelper
}

func NewGenreRepository(dbh *DataBaseHelper) GenreRepository {
	return &genreRepository{dbh}
}

// SaveGenre ..
func (repo *genreRepository) SaveGenre(genre *entities.Genre) (*entities.Genre, error) {
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
func (repo *genreRepository) GetGenre(ID int) (*entities.Genre, error) {
	mo := entities.Genre{}

	row := repo.dataBaseHelper.db.QueryRow("SELECT id, title, release_date age FROM genre where id = $1", ID)

	err := row.Scan(&mo.ID, &mo.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &mo, nil
}
