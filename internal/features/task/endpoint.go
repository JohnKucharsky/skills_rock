package task

import (
	"github.com/JohnKucharsky/skills_rock/internal/domain"
	"github.com/JohnKucharsky/skills_rock/internal/shared"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	Service interface {
		Create(ctx *fiber.Ctx) error
		Get(ctx *fiber.Ctx) error
		GetOne(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	service struct {
		repository StoreI
	}
)

func New(store *Store) Service {
	return &service{repository: store}
}

// Create a new task
// @Summary Create a new task
// @Description Create a new task, body required
// @Tags Task
// @Accept json
// @Produce json
// @Param task body domain.TaskInput true "Task object"
// @Success 201 {object} domain.Task "Task created successfully"
// @Failure 400 {object} shared.ErrorResponse "Bad request"
// @Failure 422 {object} shared.ErrorResponse "Unprocessable entity"
// @Router /task [post]
func (h *service) Create(c *fiber.Ctx) error {
	var input domain.TaskInput
	if err := shared.BindBody(c, &input); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(shared.ErrorRes(err.Error()))
	}
	one, err := h.repository.Create(c.Context(), input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(shared.ErrorRes(err.Error()))
	}
	return c.Status(http.StatusCreated).JSON(one)
}

// Get all tasks
// @Summary Get all tasks
// @Description Retrieve all tasks
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {array} domain.Task "List of tasks"
// @Failure 500 {object} shared.ErrorResponse "Internal server error"
// @Router /task [get]
func (h *service) Get(c *fiber.Ctx) error {
	task, err := h.repository.GetMany(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(shared.ErrorRes(err.Error()))
	}
	return c.Status(http.StatusOK).JSON(task)
}

// GetOne
// Get a task by ID
// @Summary Get a task by ID
// @Description Retrieve a specific task using its ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} domain.Task "Task details"
// @Failure 400 {object} shared.ErrorResponse "Bad request"
// @Failure 422 {object} shared.ErrorResponse "Unprocessable entity"
// @Router /task/{id} [get]
func (h *service) GetOne(c *fiber.Ctx) error {
	id, err := shared.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(shared.ErrorRes(err.Error()))
	}
	task, err := h.repository.GetOne(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(shared.ErrorRes(err.Error()))
	}
	return c.Status(http.StatusOK).JSON(task)
}

// Update an existing task
// @Summary Update a task
// @Description Update an existing task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body domain.TaskInput true "Updated task data"
// @Success 200 {object} domain.Task "Task updated successfully"
// @Failure 400 {object} shared.ErrorResponse "Bad request"
// @Failure 422 {object} shared.ErrorResponse "Unprocessable entity"
// @Router /task/{id} [put]
func (h *service) Update(c *fiber.Ctx) error {
	id, err := shared.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(shared.ErrorRes(err.Error()))
	}
	var req domain.TaskInput
	if err := shared.BindBody(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(shared.ErrorRes(err.Error()))
	}
	updatedEntity, err := h.repository.Update(c.Context(), req, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(shared.ErrorRes(err.Error()))
	}
	return c.Status(http.StatusOK).JSON(updatedEntity)
}

// Delete a task
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} shared.IDResponse "Task deleted successfully"
// @Failure 400 {object} shared.ErrorResponse "Bad request"
// @Failure 422 {object} shared.ErrorResponse "Unprocessable entity"
// @Router /task/{id} [delete]
func (h *service) Delete(c *fiber.Ctx) error {
	id, err := shared.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(shared.ErrorRes(err.Error()))
	}
	ID, err := h.repository.Delete(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(shared.ErrorRes(err.Error()))
	}
	return c.Status(http.StatusOK).JSON(shared.IDResponse{
		ID: ID,
	})
}
