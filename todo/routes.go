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
	group.Get("/", userController.listAllTodo)
	group.Get("/new", userController.newTodo)
	group.Post("/new", userController.saveNewTodo)
	group.Get("/:id<int>", userController.editTodo)
	group.Post("/:id<int>", userController.saveEditTodo)
	group.Get("/:id<int>/delete", userController.deleteTodo)
}
