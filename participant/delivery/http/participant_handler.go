package http

import (
	"eventzezz_backend/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"regexp"
)

type ParticipantHandler struct {
	ParticipanUsecase domain.ParticipantUsecase
}

func NewParticipantHandler(app *fiber.App, us domain.ParticipantUsecase) {
	handler := &ParticipantHandler{
		ParticipanUsecase: us,
	}

	app.Get("/participant/:eventID", handler.GetParticipantByEventID)
	app.Post("/participant", handler.CreateParticipant)
	app.Delete("/participant", handler.DeleteParticipantEvent)

	// TODO: remove this endpoint not using this endpoint
	//app.Get("/attendance/:eventID", handler.GetAttendanceByEventID)
	//app.Post("/join/event", handler.JoinEvent)
	//app.Post("/clockin", handler.ClockInEvent)
	//app.Post("/check/registrasi/event", handler.CheckParticipantRegister)

}

func (u *ParticipantHandler) GetParticipantByEventID(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	eventUUID, _ := uuid.Parse(eventID)
	responses, err := u.ParticipanUsecase.GetParticipantsByEventID(c.Context(), eventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: responses})
}

func (u *ParticipantHandler) DeleteParticipantEvent(c *fiber.Ctx) error {
	participantEventIDs := c.FormValue("participant_event_ids")
	fmt.Println(string(participantEventIDs))

	r := regexp.MustCompile(`[^\[\],]+`)
	participantEventIDSlice := r.FindAllString(string(participantEventIDs), -1)
	participantEventUUID := stringToUUIDS(participantEventIDSlice)
	err := u.ParticipanUsecase.DeleteParticipantEventByID(c.Context(), participantEventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}

func (u *ParticipantHandler) ClockInEvent(c *fiber.Ctx) error {
	joinEventID := c.FormValue("join_event_id")
	joinEventUUID, err := uuid.Parse(joinEventID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	participan, err := u.ParticipanUsecase.ClockInEvent(c.Context(), joinEventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: participan})
}

func (u *ParticipantHandler) CreateParticipant(c *fiber.Ctx) error {
	var participant domain.Participant
	err := c.BodyParser(&participant)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	err = u.ParticipanUsecase.CreateParticipant(c.Context(), &participant)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}

func (u *ParticipantHandler) JoinEvent(c *fiber.Ctx) error {
	var participant domain.Participant
	err := c.BodyParser(&participant)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})

	}
	eventID := c.FormValue("event_id")

	eventUUID, _ := uuid.Parse(eventID)

	err = u.ParticipanUsecase.JoinEvent(c.Context(), &participant, eventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}

func (u *ParticipantHandler) CheckParticipantRegister(c *fiber.Ctx) error {
	email := c.FormValue("email")
	eventID := c.FormValue("event_id")
	eventUUID, _ := uuid.Parse(eventID)

	participant, message, _ := u.ParticipanUsecase.CheckParticipantRegister(c.Context(), email, eventUUID)

	response := struct {
		Participant domain.Participant `json:"participant"`
		Message     string             `json:"message"`
	}{
		Participant: participant,
		Message:     message,
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: response})
}

func (u *ParticipantHandler) GetAttendanceByEventID(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	eventUUID, _ := uuid.Parse(eventID)

	responses, _ := u.ParticipanUsecase.GetAttendanceByEventID(c.Context(), eventUUID)
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: responses})
}
