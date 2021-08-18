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
		err = rows.Scan(&d.ID, &d.Name, &d.Stock)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, err
}

// FindById...
func (model bookRepository) FindById(id int64) (data models.Book, err error) {
	query := models.BookSelectStatement + ` WHERE id = $1 LIMIT 1`
	err = model.DB.QueryRow(query, helper.EmptyId(id)).Scan(&data.ID, &data.Name, &data.Stock)
	return data, err
}

// Add...
func (model bookRepository) Add(body *viewmodel.BookVM) (res int64, err error) {
	query := `INSERT INTO books (bookname, stock) values ($1, $2) RETURNING id`
	err = model.DB.QueryRow(query, body.Name, body.Stock).Scan(&res)

	return res, err
}

// Edit...
func (model bookRepository) Edit(id int64, body *viewmodel.BookVM) (res int64, err error) {
	query := `UPDATE books set bookname = $1, stock = $2 where id = $3 returning id`
	err = model.DB.QueryRow(query, body.Name, body.Stock, id).Scan(&res)
	return res, err
}

// Delete...
func (model bookRepository) Delete(id int64) (res int64, err error) {
	query := `DELETE FROM books where id = $1 returning id`
	err = model.DB.QueryRow(query, id).Scan(&res)
	return res, err
}
