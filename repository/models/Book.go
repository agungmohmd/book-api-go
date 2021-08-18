package models

type Book struct {
	ID    int64  `db:"id"`
	Name  string `db:"bookname"`
	Stock int    `db:"stock"`
}

var (
	BookSelectStatement = `select * from books`
	BookWhereStock      = `stock > 0`
	BookWhereID         = `id = `
)
