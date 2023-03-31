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

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(app *fiber.App, us domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: us,
	}
	app.Post("/auth/register", handler.Register)
	app.Post("/auth/login", handler.Login)
	app.Get("/auth/me", middleware.JWTMiddleware, handler.GetMe)
}

func (u *AuthHandler) Register(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	username := c.FormValue("username")
	result, err := u.AuthUsecase.Register(c.Context(), username, password, email)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *AuthHandler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	result, isSuccess, _ := u.AuthUsecase.Login(c.Context(), email, password)
	if !isSuccess {
		return c.JSON(domain.WebResponse{
			Code:   401,
			Status: "UNAUTHORIZED",
			Data:   "Email or Password is wrong",
		})
	}
	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}

func (u *AuthHandler) GetMe(c *fiber.Ctx) error {
	userID, ok := c.Context().Value("userID").(string)
	if !ok {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: "Error when parsing userID"})
	}

	userUUID, _ := uuid.Parse(userID)

	result, err := u.AuthUsecase.GetMe(c.Context(), userUUID)
	if err != nil {
		return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
	}

	return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
}
