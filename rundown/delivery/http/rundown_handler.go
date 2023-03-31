package http

import (
	"eventzezz_backend/domain"
	"eventzezz_backend/rundown/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type RundownHandler struct {
	RundownUsecase domain.RundownUsecase
}

func NewRundownHandler(app *fiber.App, us domain.RundownUsecase) {
	handler := &RundownHandler{
		RundownUsecase: us,
	}
	app.Get("/rundowns/:eventID", handler.FetchRundowns)
	app.Get("/rundown/:id", handler.GetRundownByID)
	app.Post("/rundowns", handler.CreateRundown)
	app.Put("/rundown", handler.UpdateRundown)
	app.Delete("/rundown/:id", handler.DeleteRoundown)
}

func (u *RundownHandler) FetchRundowns(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	eventUUID, _ := uuid.Parse(eventID)
	pageS := c.Query("page")
	page, _ := strconv.Atoi(pageS)
	limitS := c.Query("limit")
	limit, _ := strconv.Atoi(limitS)

	listAr, paginate, err := u.RundownUsecase.GetRundownsByEventId(c.Context(), eventUUID, page, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when fetching tasks",
			"data":    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(domain.WebResponse{
		Code:     200,
		Status:   "success",
		Data:     listAr,
		Paginate: paginate,
	})
}

func (u *RundownHandler) GetRundownByID(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, _ := uuid.Parse(id)

	rundowns, err := u.RundownUsecase.GetRundownByID(c.Context(), uuid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when fetching tasks",
			"data":    err,
		})
	}
	return c.Status(http.StatusOK).JSON(domain.WebResponse{
		Code:   200,
		Status: "success",
		Data:   rundowns,
	})
}

func (u *RundownHandler) CreateRundown(c *fiber.Ctx) error {
	var rundown domain.Rundown
	if err := c.BodyParser(&rundown); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when parsing body",
			"data":    nil,
		})
	}

	eventID := c.FormValue("event_id")
	eventUUID, _ := uuid.Parse(eventID)
	rundown.EventID = eventUUID
	start := c.FormValue("start_date")
	end := c.FormValue("end_date")
	startDate, _ := time.Parse("2006-01-02 15:04:05", start)
	endDate, _ := time.Parse("2006-01-02 15:04:05", end)

	rundown.StartDate = startDate
	rundown.EndDate = endDate
	rundown.EventID = eventUUID

	err := u.RundownUsecase.CreateRundown(c.Context(), &rundown)

	rundownResp := usecase.ToRoundownResponse(rundown)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when creating rundown",
			"data":    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(domain.WebResponse{
		Code:   200,
		Status: "success",
		Data:   rundownResp,
	})
}

func (u *RundownHandler) DeleteRoundown(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, _ := uuid.Parse(id)
	err := u.RundownUsecase.DeleteRundownByID(c.Context(), uuid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when fetching tasks",
			"data":    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(domain.WebResponse{
		Code:   200,
		Status: "success",
		Data:   nil,
	})
}

func (u *RundownHandler) UpdateRundown(c *fiber.Ctx) error {

	var rundown domain.Rundown
	if err := c.BodyParser(&rundown); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when parsing body",
			"data":    nil,
		})
	}

	eventID := c.FormValue("event_id")
	eventUUID, _ := uuid.Parse(eventID)
	rundown.EventID = eventUUID
	start := c.FormValue("start_date")
	end := c.FormValue("end_date")
	startDate, _ := time.Parse("2006-01-02 15:04:05", start)
	endDate, _ := time.Parse("2006-01-02 15:04:05", end)
	rundown.StartDate = startDate
	rundown.EndDate = endDate
	rundown.EventID = eventUUID
	err := u.RundownUsecase.UpdateRundown(c.Context(), &rundown)

	rundownResp := usecase.ToRoundownResponse(rundown)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error when creating rundown",
			"data":    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(domain.WebResponse{
		Code:   200,
		Status: "success",
		Data:   rundownResp,
	})
}
