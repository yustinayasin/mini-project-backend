package controllers

import (
	"fmt"
	"kemejaku/business/users"
	"kemejaku/controllers"
	"kemejaku/controllers/users/request"
	"kemejaku/controllers/users/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}

// dipasangkan dengan routing
func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) LoginController(c echo.Context) error {
	ctx := c.Request().Context()

	//dari request
	var userLogin request.UserLogin

	err := c.Bind(&userLogin)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if userLogin.Email == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Email empty")
	}

	if userLogin.Password == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Password empty")
	}

	//error apa ni
	user, errRepo := controller.usecase.LoginController(*userLogin.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no account with that password and email", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) GetAllUsersController(c echo.Context) error {
	ctx := c.Request().Context()

	user, errRepo := controller.usecase.GetAllUsersController(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no user column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(user))
}

func (controller *UserController) SignUpController(c echo.Context) error {
	ctx := c.Request().Context()

	var userSignup request.UserLogin

	err := c.Bind(&userSignup)

	fmt.Println(userSignup)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if userSignup.Email == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Email empty")
	}

	if userSignup.Password == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Password empty")
	}

	user, errRepo := controller.usecase.SignUpController(*userSignup.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Email and password doesn't match", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) GetDetailUserController(c echo.Context) error {
	ctx := c.Request().Context()

	userId, _ := strconv.Atoi(c.Param("userId"))

	if userId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "User ID empty")
	}

	user, errRepo := controller.usecase.GetDetailUserController(userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) EditUserController(c echo.Context) error {
	ctx := c.Request().Context()

	var userEdit request.UserEdit
	userId, _ := strconv.Atoi(c.Param("userId"))

	err := c.Bind(&userEdit)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if userEdit.Email == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Email empty")
	}

	if userEdit.Password == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Password empty")
	}

	if userId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "User ID empty")
	}

	user, errRepo := controller.usecase.EditUserController(*userEdit.ToUsecase(), userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) DeleteUserController(c echo.Context) error {
	ctx := c.Request().Context()

	userId, _ := strconv.Atoi(c.Param("userId"))

	if userId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "User ID empty")
	}

	user, errRepo := controller.usecase.DeleteUserController(userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}
