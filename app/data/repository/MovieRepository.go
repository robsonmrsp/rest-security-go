package repository

import (
	"fmt"
	"math/rand"

	"github.com/robsonmrsp/rest-security-go/app/domain/data"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

type movieRepository struct {
	dataBaseHelper *DataBaseHelper
}

func NewMovieRepository(dbh *DataBaseHelper) data.MovieRepository {
	return &movieRepository{dbh}
}

// SaveMovie ..
func (repo *movieRepository) SaveMovie(movie *entities.Movie) (*entities.Movie, error) {
	newID := rand.Intn(10000)
	stmt, err := repo.dataBaseHelper.tx.Prepare(`INSERT INTO movie(id, title, release_date) VALUES( $1, $2, $3 )`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	if _, err := stmt.Exec(newID, movie.Title, movie.ReleaseDate); err != nil {
		return nil, err
	}
	movie.ID = newID

	return movie, nil
}

// GetMovie ...
func (repo *movieRepository) GetMovie(ID int) (*entities.Movie, error) {
	mo := entities.Movie{}

	row := repo.dataBaseHelper.db.QueryRow("SELECT id, title, release_date age FROM movie where id = $1", ID)

	err := row.Scan(&mo.ID, &mo.Title, &mo.ReleaseDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &mo, nil
}

// GetPageMovie ...
func (repo *movieRepository) GetPageMovie(page int, pageSize int, order string, orderBy string, filter parser.Parameters) (*[]entities.Movie, error) {
	movies := []entities.Movie{}
	if filter.Exists("title") {

	}
	rows, err := repo.dataBaseHelper.db.Query("SELECT id, title, release_date age FROM movie where id > $1 order by id OFFSET $2 LIMIT $3", 0, (pageSize)*(page-1), pageSize)
	// We still need to create a new
	// query for count
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mo := entities.Movie{}
		err = rows.Scan(&mo.ID, &mo.Title, &mo.ReleaseDate)
		movies = append(movies, mo)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &movies, nil
}
