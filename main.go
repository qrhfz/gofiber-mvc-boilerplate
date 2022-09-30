package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/jackc/pgx/v5/pgxpool"
	"qori.dev/fiber-todo/todo"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	dbpool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/main-layout",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	todo.RegisterRoute(app, dbpool)

	app.Listen(":3000")
}
