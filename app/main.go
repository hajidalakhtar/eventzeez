package main

import (
	_participantHttpDelivery "eventzezz_backend/participant/delivery/http"
	_participantRepo "eventzezz_backend/participant/repository/posgresql"
	_participantUcase "eventzezz_backend/participant/usecase"

	_userHttpDelivery "eventzezz_backend/user/delivery/http"
	_userRepo "eventzezz_backend/user/repository/posgresql"
	_userUcase "eventzezz_backend/user/usecase"

	_authHttpDelivery "eventzezz_backend/authentication/delivery/http"
	_authRepo "eventzezz_backend/authentication/repository/posgresql"
	_authUcase "eventzezz_backend/authentication/usecase"

	_eventHttpDelivery "eventzezz_backend/event/delivery/http"
	_eventRepo "eventzezz_backend/event/repository/posgresql"
	_eventUcase "eventzezz_backend/event/usecase"

	_taskHttpDelivery "eventzezz_backend/task/delivery/http"
	_taskRepo "eventzezz_backend/task/repository/posgresql"
	_taskUcase "eventzezz_backend/task/usecase"

	_budgetHttpDelivery "eventzezz_backend/budget/delivery/http"
	_budgetRepo "eventzezz_backend/budget/repository/posgresql"
	_budgetUcase "eventzezz_backend/budget/usecase"

	_attendanceHttpDelivery "eventzezz_backend/attendance/delivery/http"
	_attendanceRepo "eventzezz_backend/attendance/repository/posgresql"
	_attendanceUcase "eventzezz_backend/attendance/usecase"

	_rundownHttpDelivery "eventzezz_backend/rundown/delivery/http"
	_rundownRepo "eventzezz_backend/rundown/repository/posgresql"
	_rundownUcase "eventzezz_backend/rundown/usecase"

	"eventzezz_backend/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"log"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)

	//viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(struct {
			Code   int    `json:"code"`
			Status string `json:"status"`
			Data   string `json:"data"`
		}{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   string `json:"data"`
	}{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}

func main() {

	// Init Config
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	portService := viper.GetString(`server.address`)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	dbConn := database.NewPosgresqlDatabase(dbHost, dbPort, dbUser, dbPass, dbName)
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	app.Use(cors.New())

	// Init Repository
	ur := _userRepo.NewPosgresqlUserRepository(dbConn)
	pr := _participantRepo.NewPosgresqlParticipantRepository(dbConn)
	er := _eventRepo.NewPosgresqlEventRepository(dbConn)
	ar := _authRepo.NewPosgresqlAuthRepository(dbConn)
	tr := _taskRepo.NewPosgresqlTaskRepository(dbConn)
	br := _budgetRepo.NewPosgresqlBudgetRepository(dbConn)
	atr := _attendanceRepo.NewPosgresqlAttendanceRepository(dbConn)
	rr := _rundownRepo.NewPosgresqlRundownRepository(dbConn)

	// Init Usecase
	uu := _userUcase.NewUserUsecase(ur, timeoutContext)
	pu := _participantUcase.NewParticipantUsecase(pr)
	eu := _eventUcase.NewEventUsecase(er, timeoutContext)
	au := _authUcase.NewAuthUsecase(ar, ur, timeoutContext)
	tu := _taskUcase.NewTaskUsecase(tr, timeoutContext)
	bu := _budgetUcase.NewBudgetUsecase(br, timeoutContext)
	atu := _attendanceUcase.NewAttendanceUsecase(atr, pr, timeoutContext)
	ru := _rundownUcase.NewRundownUsecase(rr, timeoutContext)

	// Init Delivery
	_userHttpDelivery.NewUserHandler(app, uu)
	_participantHttpDelivery.NewParticipantHandler(app, pu)
	_eventHttpDelivery.NewEventHandler(app, eu)
	_authHttpDelivery.NewAuthHandler(app, au)
	_taskHttpDelivery.NewTaskHandler(app, tu)
	_budgetHttpDelivery.NewBudgetHandler(app, bu)
	_attendanceHttpDelivery.NewAttendanceHandler(app, atu)
	_rundownHttpDelivery.NewRundownHandler(app, ru)

	// Run Service
	err := app.Listen(portService)
	if err != nil {
		panic(err)
	}
}
