package services

import (
	"errors"
	"fmt"
	"github.com/robsonmrsp/rest-security-go/models"
)

var (
	people = map[int]models.Person{
		1: {
			ID:    1,
			Name:  "Marcio",
			Age:   41,
			Email: "robsonmrsp@gmail.com",
		},
		2: {
			ID:    2,
			Name:  "Robson",
			Age:   14,
			Email: "marciomrsp@gmail.com",
		},
	}
)
// SavePerson ...
func SavePerson(person models.Person) models.Person {
	nextID := len(people) + 1
	person.ID = nextID
	people[nextID] = person
	return person
}

// GetPerson ..
func GetPerson(id int) (models.Person, error) {
	people, ok := people[id]
	if ok {
		return people, nil
	}
	return models.Person{}, errors.New("Pessoa não encontrada")
}

func main() {
	fmt.Println("só pra dizer")
}
