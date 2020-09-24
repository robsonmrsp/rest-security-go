package services

// https://www.alexedwards.net/blog/organising-database-access

import (
	"errors"
	"fmt"
	"log"

	"github.com/robsonmrsp/rest-security-go/app/data/repository"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
)

// SavePerson ...
func SavePerson(person *entities.Person) (*entities.Person, error) {
	p, err := repository.SavePerson(person)
	if err != nil {
		log.Print("Error saving person", err)
		return nil, err
	}

	return p, nil
}

// GetPerson ..
func GetPerson(id int) (*entities.Person, error) {
	person, err := repository.GetPerson(id)

	if err != nil {
		return &entities.Person{}, errors.New("Pessoa não encontrada")
	}
	return person, nil
}

func main() {
	fmt.Println("só pra dizer")
}
