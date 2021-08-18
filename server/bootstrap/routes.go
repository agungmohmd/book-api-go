package bootstrap

import (
	"github.com/agungmohmd/books-api/server/bootstrap/routes"
	"github.com/agungmohmd/books-api/server/handlers"
)

func (boot Bootstrap) RegisterRouters() {
	handler := handlers.Handler{
		FiberApp:   boot.App,
		ContractUC: &boot.ContractUC,
	}

	apiV1 := boot.App.Group("/V1")
	bookRoutes := routes.BookRoute{RouterGroup: apiV1, Handler: handler}
	bookRoutes.RegisterRoute()
}
