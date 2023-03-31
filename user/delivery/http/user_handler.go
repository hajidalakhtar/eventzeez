package http

import (
	"eventzezz_backend/domain"
	"eventzezz_backend/user/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(app *fiber.App, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}

	app.Get("/users", middleware.JWTMiddleware, handler.FetchUsers)
	app.Get("/user/:id", middleware.JWTMiddleware, handler.GetUserByID)
	app.Post("/user", middleware.JWTMiddleware, handler.CreateUser)
	app.Post("/user/:id", middleware.JWTMiddleware, handler.UpdateUser)
	app.Delete("/user/:id", middleware.JWTMiddleware, handler.DeleteUser)
	app.Get("/user/email/:email", middleware.JWTMiddleware, handler.GetUserByEmail)
}

func (u *UserHandler) FetchUsers(c *fiber.Ctx) error {
	result, _, err := u.UserUsecase.GetUsers(c.Context(), "0", 0)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	userUUID, _ := uuid.Parse(userID)

	result, err := u.UserUsecase.GetUserByID(c.Context(), userUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	userUUID, _ := uuid.Parse(userID)

	var user domain.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	err = u.UserUsecase.UpdateUser(c.Context(), &user, userUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})
}

func (u *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	userUUID, _ := uuid.Parse(userID)

	err := u.UserUsecase.DeleteUserByID(c.Context(), userUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})

	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})

}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})

	}
	err = u.UserUsecase.CreateUser(c.Context(), &user)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: "Success"})

}

func (u *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	return c.JSON("Get User By Email")
}
