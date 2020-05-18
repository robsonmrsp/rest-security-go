package entities

type Person struct {
	ID    int    `db:"id" `
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Email string `db:"email"`
}
