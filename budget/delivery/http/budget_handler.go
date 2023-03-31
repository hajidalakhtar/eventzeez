package http

import (
	"eventzezz_backend/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type BudgetHandler struct {
	BudgetUsecase domain.BudgetUsecase
}

func NewBudgetHandler(app *fiber.App, us domain.BudgetUsecase) {
	handler := &BudgetHandler{
		BudgetUsecase: us,
	}
	app.Get("/budget/:eventID", handler.GetBudgetsByEventID)
	app.Get("/budget/detail/:id", handler.GetBudgetByID)
	app.Post("/budget", handler.CreateBudget)
	app.Delete("/budget/:id", handler.DeleteBudget)
}

func (u *BudgetHandler) GetBudgetsByEventID(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	eventUUID, _ := uuid.Parse(eventID)
	result, _, err := u.BudgetUsecase.GetBudgetsByEventID(c.Context(), "0", 0, eventUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *BudgetHandler) GetBudgetByID(c *fiber.Ctx) error {
	budgetID := c.Params("id")
	budgetUUID, _ := uuid.Parse(budgetID)
	result, err := u.BudgetUsecase.GetBudgetByID(c.Context(), budgetUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *BudgetHandler) CreateBudget(c *fiber.Ctx) error {
	var budget domain.Budget
	err := c.BodyParser(&budget)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	err = u.BudgetUsecase.CreateBudget(c.Context(), &budget)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}

func (u *BudgetHandler) DeleteBudget(c *fiber.Ctx) error {
	budgetID := c.Params("id")
	budgetUUID, _ := uuid.Parse(budgetID)
	err := u.BudgetUsecase.DeleteBudgetByID(c.Context(), budgetUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}
