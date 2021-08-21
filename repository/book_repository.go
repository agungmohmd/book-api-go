package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/agungmohmd/books-api/helper"
	"github.com/agungmohmd/books-api/repository/models"
	"github.com/agungmohmd/books-api/usecase/viewmodel"
)

// IBook...
type IBook interface {
	SelectAll(search string) ([]models.Book, error)
	FindAll(parameter models.BookParameter) ([]models.Book, int, error)
	FindById(id int64) (models.Book, error)
	Add(body *viewmodel.BookVM) (int64, error)
	Edit(id int64, body *viewmodel.BookVM) (int64, error)
	Delete(id int64) (int64, error)
}

// bookRepository...
type bookRepository struct {
	DB *sql.DB
}

// NewBookRepository...
func NewBookRepository(db *sql.DB) IBook {
	return &bookRepository{DB: db}
}

// Scan Rows
func (model bookRepository) ScanRows(rows *sql.Rows) (res models.Book, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.Stock,
	)

	if err != nil {
		return res, err
	}
	return res, nil
}

// Scan Row
// func (model bookRepository) scanRow(row *sql.Row) (res models.Book, err error) {
// 	err = row.Scan(
// 		&res.ID, &res.Name, &res.Stock,
// 	)
// 	if err != nil {
// 		return res, err
// 	}

// 	return res, nil
// }

// SelectAll...
func (model bookRepository) SelectAll(search string) (data []models.Book, err error) {
	fmt.Println("the function is used")
	query := models.BookSelectStatement + ` Where bookname Like $1`
	fmt.Println(query)
	rows, err := model.DB.Query(query, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		d := models.Book{}
		err = rows.Scan(&d.ID, &d.Name, &d.Stock, &d.UpdatedAt)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, err
}

// FindAll...
func (model bookRepository) FindAll(parameter models.BookParameter) (data []models.Book, count int, err error) {
	conditionString := ``
	query := models.BookSelectStatement + ` WHERE ` + models.BookWhereStock + ` AND (LOWER(def."bookname") LIKE $1) ` + conditionString + ` ORDER BY  ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	// fmt.Println(query)
	// fmt.Println(`by : ` + parameter.By)
	// fmt.Println(`sort : ` + parameter.Sort)
	// fmt.Println(`search : ` + parameter.Search)
	// fmt.Println(`offset : ` + strconv.Itoa(parameter.Offset))
	// fmt.Println(`limit : ` + strconv.Itoa(parameter.Limit))
	rows, err := model.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}
	defer rows.Close()
	// for rows.Next() {
	// 	temp, err := model.ScanRows(rows)
	// 	if err != nil {
	// 		return data, count, err
	// 	}
	// 	data = append(data, temp)
	// }
	for rows.Next() {
		d := models.Book{}
		err = rows.Scan(&d.ID, &d.Name, &d.Stock, &d.UpdatedAt)
		// println(err)
		if err != nil {
			return data, 0, err
		}
		data = append(data, d)
	}
	// fmt.Println(data)
	err = rows.Err()
	if err != nil {
		return data, count, err
	}

	query = `SELECT COUNT(*) FROM books def where ` + models.BookWhereStock + ` AND (LOWER(def."bookname") LIKE $1) ` + conditionString
	err = model.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindById...
func (model bookRepository) FindById(id int64) (data models.Book, err error) {
	query := models.BookSelectStatement + ` WHERE id = $1 LIMIT 1`
	err = model.DB.QueryRow(query, helper.EmptyId(id)).Scan(&data.ID, &data.Name, &data.Stock, &data.UpdatedAt)
	return data, err
}

// Add...
func (model bookRepository) Add(body *viewmodel.BookVM) (res int64, err error) {
	query := `INSERT INTO books (bookname, stock, updated_at) values ($1, $2, $3) RETURNING id`
	err = model.DB.QueryRow(query, body.Name, body.Stock, body.UpdatedAt).Scan(&res)

	return res, err
}

// Edit...
func (model bookRepository) Edit(id int64, body *viewmodel.BookVM) (res int64, err error) {
	query := `UPDATE books set bookname = $1, stock = $2 where id = $3 returning id`
	err = model.DB.QueryRow(query, body.Name, body.Stock, body.UpdatedAt, id).Scan(&res)
	return res, err
}

// Delete...
func (model bookRepository) Delete(id int64) (res int64, err error) {
	query := `DELETE FROM books where id = $1 returning id`
	err = model.DB.QueryRow(query, id).Scan(&res)
	return res, err
}
