package entities

type Genre struct {
	ID   int    `db:"id" `
	name string `db:"name"`
}
