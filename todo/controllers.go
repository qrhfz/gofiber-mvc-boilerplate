package todo

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
	DbPool *pgxpool.Pool
}

func (uc *UserController) listAllTodo(ctx *fiber.Ctx) error {
	rows, err := uc.DbPool.Query(context.Background(), "SELECT * FROM todo")

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	todos := []TodoModel{}

	for rows.Next() {

		var todo = TodoModel{}

		rows.Scan(&todo.Id, &todo.Todo, &todo.Done)

		todos = append(todos, todo)
	}

	return ctx.Render("views/index", todos)
}

func (uc *UserController) addTodo(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"aa": "aaa",
	})
}

func (uc *UserController) getTodo(ctx *fiber.Ctx) error {
	ctx.ParamsInt("id")
	return ctx.JSON(fiber.Map{
		"aa": "aaa",
	})
}
func (uc *UserController) editTodo(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"aa": "aaa",
	})
}

func (uc *UserController) deleteTodo(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"aa": "aaa",
	})
}
