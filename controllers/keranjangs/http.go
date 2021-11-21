package controllers

import (
	"kemejaku/business/keranjangs"
	"kemejaku/controllers"
	"kemejaku/controllers/keranjangs/request"
	"kemejaku/controllers/keranjangs/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KeranjangController struct {
	usecase keranjangs.KeranjangUseCaseInterface
}

func NewKeranjangController(uc keranjangs.KeranjangUseCaseInterface) *KeranjangController {
	return &KeranjangController{
		usecase: uc,
	}
}

func (controller *KeranjangController) InsertKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var keranjang request.KeranjangInsert

	err := c.Bind(&keranjang)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if keranjang.IdUser == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "User ID empty")
	}

	user, errRepo := controller.usecase.InsertKeranjang(*keranjang.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Failed to insert new kemeja keranjang", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *KeranjangController) GetAllKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	keranjang, errRepo := controller.usecase.GetAllKeranjang(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no keranjang column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(keranjang))
}

func (controller *KeranjangController) GetKeranjangDetail(c echo.Context) error {
	ctx := c.Request().Context()

	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	if keranjangId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja keranjang ID empty")
	}

	keranjang, errRepo := controller.usecase.GetKeranjangDetail(keranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(keranjang))
}

func (controller *KeranjangController) EditKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var keranjang request.KeranjangEdit
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))
	err := c.Bind(&keranjang)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if keranjangId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja keranjang ID empty")
	}

	keranjangRepo, errRepo := controller.usecase.EditKeranjang(*keranjang.ToUsecase(), keranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(keranjangRepo))
}

func (controller *KeranjangController) DeleteKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	if keranjangId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Keranjang ID empty")
	}

	keranjang, errRepo := controller.usecase.DeleteKeranjang(keranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(keranjang))
}