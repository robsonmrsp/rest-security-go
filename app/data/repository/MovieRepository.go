package repository

import (
	"fmt"
	"math/rand"

	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

type MovieRepository struct {
	dataBaseHelper *DataBaseHelper
}

func NewMovieRepository(dbh *DataBaseHelper) *MovieRepository {
	repo := new(MovieRepository)
	repo.dataBaseHelper = dbh
	return repo
}

// SaveMovie ..
func (repo *MovieRepository) SaveMovie(movie *entities.Movie) (*entities.Movie, error) {
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
func (repo *MovieRepository) GetMovie(ID int) (*entities.Movie, error) {
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
func (repo *MovieRepository) GetPageMovie(page int, pageSize int, order string, orderBy string, filter parser.Parameters) (*[]entities.Movie, error) {
	movies := []entities.Movie{}
	if filter.Exists("title") {

	}
	rows, err := repo.dataBaseHelper.db.Query("SELECT id, title, release_date age FROM movie where id > $1 order by id OFFSET $2 LIMIT $3", 1, (pageSize)*(page-1), pageSize)
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

// AllMovies ..
func AllMovies() ([]*entities.Movie, error) {
	dataBaseHelper := GetDb()
	rows, err := dataBaseHelper.db.Query("SELECT title, release_date FROM movie")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]*entities.Movie, 0)
	for rows.Next() {
		movie := new(entities.Movie)
		err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}
		people = append(people, movie)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return people, nil
}
