package repository

import (
	"fmt"
	"math/rand"

	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
)

// SavePerson ..
func SavePerson(person *entities.Person) (*entities.Person, error) {
	newID := rand.Intn(10000)
	stmt, err := dataBaseHelper.tx.Prepare(`INSERT INTO person(id, name, email, age) VALUES( $1, $2, $3, $4 )`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	if _, err := stmt.Exec(newID, person.Name, person.Email, person.Age); err != nil {
		return nil, err
	}
	person.ID = newID

	return person, nil
}

// GetPerson ...
func GetPerson(ID int) (*entities.Person, error) {
	p := entities.Person{}

	row := dataBaseHelper.db.QueryRow("SELECT id, name, email, age FROM person where id = $1", ID)

	err := row.Scan(&p.ID, &p.Name, &p.Email, &p.Age)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &p, nil
}

// AllPersons ..
func AllPersons() ([]*entities.Person, error) {
	dataBaseHelper := GetDb()
	rows, err := dataBaseHelper.db.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]*entities.Person, 0)
	for rows.Next() {
		person := new(entities.Person)
		err := rows.Scan(&person.ID, &person.Name, &person.Email, &person.Age)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return people, nil
}
