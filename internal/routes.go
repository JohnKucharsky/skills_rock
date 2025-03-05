package internal

import (
	_ "github.com/JohnKucharsky/skills_rock/cmd/docs"
	"github.com/JohnKucharsky/skills_rock/internal/features/task"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(r *fiber.App, db *pgxpool.Pool) {
	v1 := r.Group("/api")

	store := task.NewTaskStore(db)
	handler := task.New(store)

	group := v1.Group("/task")
	group.Post("/", handler.Create)
	group.Get("/", handler.Get)
	group.Get("/:id", handler.GetOne)
	group.Put("/:id", handler.Update)
	group.Delete("/:id", handler.Delete)
}
