package task

import (
	"github.com/JohnKucharsky/skills_rock/internal/domain"
	"github.com/JohnKucharsky/skills_rock/internal/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"
)

type (
	StoreI interface {
		Create(ctx *fasthttp.RequestCtx, m domain.TaskInput) (*domain.Task, error)
		GetMany(ctx *fasthttp.RequestCtx) ([]*domain.Task, error)
		GetOne(ctx *fasthttp.RequestCtx, id int) (*domain.Task, error)
		Update(ctx *fasthttp.RequestCtx, m domain.TaskInput, id int) (*domain.Task, error)
		Delete(ctx *fasthttp.RequestCtx, id int) (*int, error)
	}

	Store struct {
		db *pgxpool.Pool
	}
	IdRes struct {
		ID int `db:"id"`
	}
)

func NewTaskStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) Create(ctx *fasthttp.RequestCtx, m domain.TaskInput) (
	*domain.Task,
	error,
) {
	sql := `INSERT INTO task (title, description, status, created_at, updated_at)
        VALUES (@title, @description, @status, now(), now())
        RETURNING id, title, description, status, created_at, updated_at`
	args := pgx.NamedArgs{
		"title":       m.Title,
		"description": m.Description,
		"status":      m.Status,
	}

	return shared.GetOneRow[domain.Task](ctx, store.db, sql, args)

}

func (store *Store) GetMany(ctx *fasthttp.RequestCtx) ([]*domain.Task, error) {
	var sql = `SELECT * from task`

	return shared.GetManyRows[domain.Task](ctx, store.db, sql)
}

func (store *Store) GetOne(ctx *fasthttp.RequestCtx, id int) (*domain.Task, error) {
	sql := `SELECT * FROM task WHERE id = @id`
	args := pgx.NamedArgs{"id": id}

	return shared.GetOneRow[domain.Task](ctx, store.db, sql, args)
}

func (store *Store) Update(ctx *fasthttp.RequestCtx, m domain.TaskInput, id int) (*domain.Task, error) {
	sql := `UPDATE task SET 
			title = @title,
			description = @description,
			status = @status,
			updated_at = now()
            WHERE id = @id 
            RETURNING id, title, description, status, created_at, updated_at`
	args := pgx.NamedArgs{
		"id":          id,
		"title":       m.Title,
		"description": m.Description,
		"status":      m.Status,
	}

	return shared.GetOneRow[domain.Task](ctx, store.db, sql, args)
}

func (store *Store) Delete(ctx *fasthttp.RequestCtx, id int) (*int, error) {
	sql := `DELETE FROM task WHERE id = @id RETURNING id`
	args := pgx.NamedArgs{"id": id}

	one, err := shared.GetOneRow[IdRes](ctx, store.db, sql, args)
	if err != nil {
		return nil, err
	}

	return &one.ID, nil
}
