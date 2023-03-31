package http

import (
	"eventzezz_backend/domain"
	"eventzezz_backend/event/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ResponseError struct {
	Message string `json:"message"`
}

type EventHandler struct {
	EventUsecase domain.EventUsecase
}

func NewEventHandler(app *fiber.App, us domain.EventUsecase) {
	handler := &EventHandler{
		EventUsecase: us,
	}
	app.Get("/events", middleware.JWTMiddleware, handler.FetchEvents)
	app.Get("/event/:id", handler.GetEventByID)
	app.Post("/event", handler.CreateEvent) //TODO: BELUM JALAN
	app.Post("/event/:id", handler.UpdateEvent)
	app.Delete("/event/:id", handler.DeleteEvent)
}

func (u *EventHandler) FetchEvents(c *fiber.Ctx) error {
	userID, ok := c.Context().Value("userID").(string)
	if !ok {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: "Error when parsing userID"})
	}

	userUUID, _ := uuid.Parse(userID)

	result, _, err := u.EventUsecase.GetEvents(c.Context(), "0", 0, userUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *EventHandler) GetEventByID(c *fiber.Ctx) error {
	eventID := c.Params("id")
	eventUUID, _ := uuid.Parse(eventID)

	result, err := u.EventUsecase.GetEventByID(c.Context(), eventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *EventHandler) UpdateEvent(c *fiber.Ctx) error {
	eventID := c.Params("id")
	eventUUID, _ := uuid.Parse(eventID)
	var event domain.Event
	err := c.BodyParser(&event)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	err = u.EventUsecase.UpdateEvent(c.Context(), &event, eventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}

func (u *EventHandler) DeleteEvent(c *fiber.Ctx) error {
	eventID := c.Params("id")
	eventUUID, _ := uuid.Parse(eventID)

	err := u.EventUsecase.DeleteEventByID(c.Context(), eventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})

	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})

}

func (u *EventHandler) CreateEvent(c *fiber.Ctx) error {
	var event domain.Event
	err := c.BodyParser(&event)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})

	}
	eventName := c.FormValue("event_name")
	eventTime := c.FormValue("event_time")
	eventDate := c.FormValue("event_date")
	eventDateFormat, _ := time.Parse("2006-01-02", eventDate)

	authorID := c.FormValue("author_id")
	authorUUID, _ := uuid.Parse(authorID)

	event.EventName = eventName
	event.EventTime = eventTime
	event.AuthorID = authorUUID
	event.EventDate = eventDateFormat

	result, err := u.EventUsecase.CreateEvent(c.Context(), &event)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}
