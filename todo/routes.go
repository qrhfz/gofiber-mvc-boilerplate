package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoute(app *fiber.App, dbpool *pgxpool.Pool) {
	userController := UserController{
		DbPool: dbpool,
	}

	group := app.Group("/")
	group.Post("/", userController.addTodo)
	group.Get("/", userController.listAllTodo)
	group.Get("/:id<int>", userController.getTodo)
	group.Get("/:id<int>/edit", userController.editTodo)
	group.Get("/:id<int>/delete", userController.deleteTodo)
}
