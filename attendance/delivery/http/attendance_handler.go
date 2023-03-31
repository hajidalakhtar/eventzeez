package http

import (
	"eventzezz_backend/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type AttendanceHandler struct {
	AttendanceUsecase domain.AttendanceUsecase
}

func NewAttendanceHandler(app *fiber.App, us domain.AttendanceUsecase) {
	handler := &AttendanceHandler{
		AttendanceUsecase: us,
	}
	app.Get("/attendance/:eventID", handler.GetAttendanceByEventID)
	app.Post("/join/event", handler.JoinEvent)
	app.Post("/clockin", handler.ClockInEvent)
	app.Post("/check/registrasi/event", handler.CheckParticipantRegister)
}

func (u *AttendanceHandler) JoinEvent(c *fiber.Ctx) error {
	var participant domain.Participant
	err := c.BodyParser(&participant)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})

	}
	eventID := c.FormValue("event_id")

	eventUUID, _ := uuid.Parse(eventID)

	err = u.AttendanceUsecase.JoinEvent(c.Context(), &participant, eventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}

func (u *AttendanceHandler) CheckParticipantRegister(c *fiber.Ctx) error {
	email := c.FormValue("email")
	eventID := c.FormValue("event_id")
	eventUUID, _ := uuid.Parse(eventID)

	participant, message, _ := u.AttendanceUsecase.CheckParticipantRegister(c.Context(), email, eventUUID)

	response := struct {
		Participant domain.Participant `json:"participant"`
		Message     string             `json:"message"`
	}{
		Participant: participant,
		Message:     message,
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: response})
}

func (u *AttendanceHandler) GetAttendanceByEventID(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	eventUUID, _ := uuid.Parse(eventID)

	responses, _ := u.AttendanceUsecase.GetAttendanceByEventID(c.Context(), eventUUID)
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: responses})
}

func (u *AttendanceHandler) ClockInEvent(c *fiber.Ctx) error {
	joinEventID := c.FormValue("join_event_id")
	joinEventUUID, err := uuid.Parse(joinEventID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	participan, err := u.AttendanceUsecase.ClockInEvent(c.Context(), joinEventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: participan})
}
