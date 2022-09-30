package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoute(app *fiber.App, dbpool *pgxpool.Pool) {
	todoController := TodoController{
		DbPool: dbpool,
	}

	group := app.Group("/")
	group.Get("/", todoController.listAllTodo)
	group.Get("/new", todoController.newTodo)
	group.Post("/new", todoController.saveNewTodo)
	group.Get("/:id<int>", todoController.editTodo)
	group.Post("/:id<int>", todoController.saveEditTodo)
	group.Get("/:id<int>/delete", todoController.deleteTodo)
}
