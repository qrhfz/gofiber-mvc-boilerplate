package todo

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
	DbPool *pgxpool.Pool
}

func (uc *UserController) listAllTodo(ctx *fiber.Ctx) error {
	rows, err := uc.DbPool.Query(context.Background(), "SELECT * FROM todo")

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	todos := []TodoModel{}

	for rows.Next() {
		var todo = TodoModel{}

		err = rows.Scan(&todo.Id, &todo.Todo, &todo.Done)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		todos = append(todos, todo)
	}

	return ctx.Render("views/todo/index", todos)
}

func (uc *UserController) newTodo(ctx *fiber.Ctx) error {

	return ctx.Render("views/todo/new-todo-form", nil)
}

func (uc *UserController) saveNewTodo(ctx *fiber.Ctx) error {
	todo := ctx.FormValue("todo")
	statement := "insert into todo(todo, done) values($1, FALSE)"
	_, err := uc.DbPool.Exec(context.Background(), statement, todo)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return ctx.Redirect("/")
}

func (uc *UserController) editTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	statement := "select id, todo, done from todo where id = $1"

	todo := TodoModel{}
	row := uc.DbPool.QueryRow(context.Background(), statement, id)

	err = row.Scan(&todo.Id, &todo.Todo, &todo.Done)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return ctx.Render("views/todo/edit-todo-form", todo)

}
func (uc *UserController) saveEditTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Bad Request")

	}

	statement := "update todo set todo=$2, done=$3 where id = $1"

	todo := ctx.FormValue("todo")
	formDone := ctx.FormValue("done")

	var done bool

	if formDone == "done" {
		done = true
	}

	_, err = uc.DbPool.Exec(context.Background(), statement, id, todo, done)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return ctx.Redirect("/")
}

func (uc *UserController) deleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	statement := "delete from todo where id = $1"

	_, err = uc.DbPool.Exec(context.Background(), statement, id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return ctx.Redirect("/")

}
