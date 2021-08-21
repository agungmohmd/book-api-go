package routes

import (
	"github.com/agungmohmd/books-api/server/handlers"
	"github.com/gofiber/fiber/v2"
)

// BookRoute...
type BookRoute struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

func (route BookRoute) RegisterRoute() {
	handler := handlers.BookHandler{Handler: route.Handler}
	r := route.RouterGroup.Group("/api/book")

	r.Get("", handler.SelectAll)
	r.Get("findsome", handler.FindAll)
	r.Get("/id/:id", handler.FindById)
	r.Post("", handler.Add)
	r.Put("/id/:id", handler.Edit)
	r.Delete("/id/:id", handler.Delete)
}
