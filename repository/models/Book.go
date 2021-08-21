package models

type Book struct {
	ID        int64  `db:"id"`
	Name      string `db:"bookname"`
	Stock     int    `db:"stock"`
	UpdatedAt string `db:"updated_at"`
}

type BookParameter struct {
	ID        int64  `db:"id"`
	Name      string `db:"bookname"`
	Stock     int    `db:"stock"`
	UpdatedAt string `db:"updated_at"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
}

var (
	BookSelectStatement = `select * from books def`
	BookOrderBy         = []string{`def.id`}
	BookOrderByString   = []string{`def.id`}
	BookWhereStock      = `def.stock > 0`
	BookWhereID         = `def.id = `
)
