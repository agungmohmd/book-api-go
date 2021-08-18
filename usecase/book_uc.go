package usecase

import (
	"errors"

	"github.com/agungmohmd/books-api/repository"
	"github.com/agungmohmd/books-api/repository/models"
	request "github.com/agungmohmd/books-api/server/requests"
	"github.com/agungmohmd/books-api/usecase/viewmodel"
)

// BookUC
type BookUC struct {
	*ContractUC
}

// BuildBody...

func (uc BookUC) BuildBody(data *models.Book, res *viewmodel.BookVM) {
	res.ID = data.ID
	res.Name = data.Name
	res.Stock = int64(data.Stock)
}

// SelectAll ...
func (uc BookUC) SelectAll(search string) (res []viewmodel.BookVM, err error) {
	repo := repository.NewBookRepository(uc.DB)
	println("the funcion in uc is called")
	data, err := repo.SelectAll(search)
	if err != nil {
		return res, err
	}

	for _, r := range data {
		tmp := viewmodel.BookVM{}
		uc.BuildBody(&r, &tmp)
		res = append(res, tmp)
	}

	return res, err

}

// FindById...
func (uc BookUC) FindById(id int64) (res viewmodel.BookVM, err error) {
	repo := repository.NewBookRepository(uc.DB)
	data, err := repo.FindById(id)
	if err != nil {
		return res, errors.New("record doenst exist")
	}
	uc.BuildBody(&data, &res)
	return res, err
}

// Add ...
func (uc BookUC) Add(req *request.BookRequest) (res viewmodel.BookVM, err error) {
	repo := repository.NewBookRepository(uc.DB)
	res = viewmodel.BookVM{
		ID:    req.ID,
		Name:  req.Name,
		Stock: req.Stock,
	}
	res.ID, err = repo.Add(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Edit...
func (uc BookUC) Edit(id int64, req *request.BookRequest) (res viewmodel.BookVM, err error) {
	repo := repository.NewBookRepository(uc.DB)
	res = viewmodel.BookVM{
		ID:    1,
		Name:  req.Name,
		Stock: req.Stock,
	}
	res.ID, err = repo.Edit(id, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

// Delete...
func (uc BookUC) Delete(id int64) (res viewmodel.BookVM, err error) {
	repo := repository.NewBookRepository(uc.DB)
	res.ID, err = repo.Delete(id)
	if err != nil {
		return res, err
	}
	return res, err
}
