package http

import (
	"eventzezz_backend/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type TaskHandler struct {
	TaskUsecase domain.TaskUsecase
}

func NewTaskHandler(app *fiber.App, us domain.TaskUsecase) {
	handler := &TaskHandler{
		TaskUsecase: us,
	}
	app.Get("/tasks", handler.FetchTasks)
	app.Get("/task/:taskid", handler.GetTaskById)
	app.Post("/task", handler.CreateTask)
	app.Delete("/task/:taskid", handler.DeleteTask)
}

func (t TaskHandler) GetTaskById(c *fiber.Ctx) error {
	taskID, err := uuid.Parse(c.Params("taskid"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{Code: http.StatusInternalServerError, Status: "Error", Data: err.Error()})
	}

	task, err := t.TaskUsecase.GetTaskByID(c.Context(), taskID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{Code: http.StatusInternalServerError, Status: "Error", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: task})
}
func (t TaskHandler) FetchTasks(c *fiber.Ctx) error {
	tasks, _, err := t.TaskUsecase.GetTasks(c.Context(), "0", 0)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when fetching tasks",
			"data":    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Tasks successfully fetched",
		"data":    tasks,
	})
}
func (t TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task domain.Task
	err := c.BodyParser(&task)
	fmt.Println(task.EventID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{Code: http.StatusInternalServerError,
			Status: "Error", Data: err.Error()},
		)
	}

	err = t.TaskUsecase.CreateTask(c.Context(), &task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{Code: http.StatusInternalServerError,
			Status: "Error", Data: err.Error()},
		)
	}
	return c.Status(http.StatusOK).JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "success"})
}
func (t TaskHandler) DeleteTask(c *fiber.Ctx) error {
	taskID, err := uuid.Parse(c.Params("taskid"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{Code: http.StatusInternalServerError, Status: "Error", Data: err.Error()})
	}

	err = t.TaskUsecase.DeleteTaskByID(c.Context(), taskID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{Code: http.StatusInternalServerError, Status: "Error", Data: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "success"})
}
