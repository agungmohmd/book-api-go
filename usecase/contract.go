package usecase

import (
	"database/sql"

	"github.com/agungmohmd/books-api/pkg/str"
	"github.com/agungmohmd/books-api/usecase/viewmodel"
)

var (
	defaultLimit    = 10
	maxLimit        = 50
	defaultSort     = "asc"
	sortWhitelist   = []string{"asc", "desc"}
	passwordLength  = 6
	defaultLastPage = 0
)

type ContractUC struct {
	ReqID     string
	UserID    string
	EnvConfig map[string]string
	DB        *sql.DB
	TX        *sql.DB
}

func (uc ContractUC) setPaginationParameter(page, limit int, orderBy, sort string, orderByWhiteLists, orderByStringWhiteLists []string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	orderBy = uc.checkWhiteList(orderBy, orderByWhiteLists)
	if str.Contains(orderByStringWhiteLists, orderBy) {
		orderBy = `LOWER(` + orderBy + `)`
	}

	if !str.Contains(sortWhitelist, sort) {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, orderBy, sort
}

func (uc ContractUC) checkWhiteList(orderBy string, whiteLists []string) string {
	for _, whiteList := range whiteLists {
		if orderBy == whiteList {
			return orderBy
		}
	}

	return "updated_at"
}

func (uc ContractUC) setPaginationResponse(page, limit, total int) (paginationResponse viewmodel.PaginationVM) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	paginationResponse = viewmodel.PaginationVM{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}

	return paginationResponse
}
