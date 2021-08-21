package handlers

import (
	// "context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/agungmohmd/books-api/pkg/str"
	"github.com/agungmohmd/books-api/repository/models"
	request "github.com/agungmohmd/books-api/server/requests"
	"github.com/agungmohmd/books-api/usecase"
	"github.com/gofiber/fiber/v2"
)

// BookHandler...

type BookHandler struct {
	Handler
}

// SelectAll...
func (h *BookHandler) SelectAll(ctx *fiber.Ctx) error {
	keyword := ctx.Query("keyword")
	println("the function in handler is called")
	bookUc := usecase.BookUC{ContractUC: h.ContractUC}
	res, err := bookUc.SelectAll(keyword)
	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindAll...
func (h *BookHandler) FindAll(ctx *fiber.Ctx) error {
	// c := ctx.Locals("ctx").(context.Context)
	parameter := models.BookParameter{
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	fmt.Println(parameter.Limit)
	uc := usecase.BookUC{ContractUC: h.ContractUC}
	res, p, err := uc.FindAll(parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}
	return h.SendResponse(ctx, res, p, err, 0)
}

// FindById...
func (h *BookHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	convert, err := strconv.Atoi(id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	bookUc := usecase.BookUC{ContractUC: h.ContractUC}
	res, err := bookUc.FindById(int64(convert))
	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add...
func (h *BookHandler) Add(ctx *fiber.Ctx) error {
	input := new(request.BookRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	bookUc := usecase.BookUC{ContractUC: h.ContractUC}
	res, err := bookUc.Add(input)
	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit...
func (h *BookHandler) Edit(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, nil, http.StatusBadRequest)
	}

	convert, err := strconv.Atoi(id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	input := new(request.BookRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	bookUc := usecase.BookUC{ContractUC: h.ContractUC}
	res, err := bookUc.Edit(int64(convert), input)
	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *BookHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, nil, http.StatusBadRequest)
	}

	convert, err := strconv.Atoi(id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	bookUc := usecase.BookUC{ContractUC: h.ContractUC}
	res, err := bookUc.Delete(int64(convert))
	return h.SendResponse(ctx, res, nil, err, 0)
}
