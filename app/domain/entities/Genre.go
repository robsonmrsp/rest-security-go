package entities

type Genre struct {
	ID   int    `db:"id" `
	Name string `db:"name"`
}
