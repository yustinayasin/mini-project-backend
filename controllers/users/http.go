package controllers

import (
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

func (controller *UserController) Login(c echo.Context) error {
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
	user, errRepo := controller.usecase.Login(*userLogin.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no account with that password and email", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	user, errRepo := controller.usecase.GetAllUsers(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no user column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(user))
}

func (controller *UserController) SignUp(c echo.Context) error {
	ctx := c.Request().Context()

	var userSignup request.UserLogin

	err := c.Bind(&userSignup)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if userSignup.Email == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Email empty")
	}

	if userSignup.Password == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Password empty")
	}

	user, errRepo := controller.usecase.SignUp(*userSignup.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Error in repo", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) GetUserDetail(c echo.Context) error {
	ctx := c.Request().Context()

	userId, _ := strconv.Atoi(c.Param("userId"))

	if userId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "User ID empty")
	}

	user, errRepo := controller.usecase.GetUserDetail(userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) EditUser(c echo.Context) error {
	ctx := c.Request().Context()

	var userEdit request.UserEdit
	userId, _ := strconv.Atoi(c.Param("userId"))

	err := c.Bind(&userEdit)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if userId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "User ID empty")
	}

	if userEdit.Email == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Email empty")
	}

	if userEdit.Password == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Password empty")
	}

	user, errRepo := controller.usecase.EditUser(*userEdit.ToUsecase(), userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	userId, _ := strconv.Atoi(c.Param("userId"))

	if userId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "User ID empty")
	}

	user, errRepo := controller.usecase.DeleteUser(userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}
