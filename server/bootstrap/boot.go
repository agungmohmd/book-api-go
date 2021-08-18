package bootstrap

import (
	"database/sql"

	"github.com/agungmohmd/books-api/usecase"
	"github.com/gofiber/fiber/v2"
)

type Bootstrap struct {
	DB         *sql.DB
	App        *fiber.App
	ContractUC usecase.ContractUC
}
