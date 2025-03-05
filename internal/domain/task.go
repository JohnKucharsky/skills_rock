package domain

import "time"

type (
	Task struct {
		ID    int    `json:"id" db:"id"`
		Title string `json:"title" db:"title"`
		// @nullable
		Description *string   `json:"description" db:"description"`
		Status      string    `json:"status" db:"status"`
		CreatedAt   time.Time `json:"created_at" db:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	}

	TaskInput struct {
		Title string `json:"title" validate:"required,max=255"`
		// @nullable
		Description *string `json:"description" validate:"max=255"`
		Status      string  `json:"status" validate:"required,oneof=new in_progress done"`
	}
)
