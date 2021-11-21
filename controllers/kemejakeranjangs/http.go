package controllers

import (
	"kemejaku/business/kemejakeranjangs"
	"kemejaku/controllers"
	"kemejaku/controllers/kemejakeranjangs/request"
	"kemejaku/controllers/kemejakeranjangs/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KemejaKeranjangController struct {
	usecase kemejakeranjangs.KemejaKeranjangUseCaseInterface
}

func NewKemejaKeranjangController(uc kemejakeranjangs.KemejaKeranjangUseCaseInterface) *KemejaKeranjangController {
	return &KemejaKeranjangController{
		usecase: uc,
	}
}

func (controller *KemejaKeranjangController) InsertKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var kk request.KemejaKeranjangInsert

	err := c.Bind(&kk)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if kk.IdKemeja == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja ID empty")
	}

	if kk.IdKeranjang == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Keranjang ID empty")
	}

	if kk.Jumlah == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Jumlah empty")
	}

	if kk.Size == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Size empty")
	}

	user, errRepo := controller.usecase.InsertKemejaKeranjang(*kk.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Email and password doesn't match", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *KemejaKeranjangController) GetAllKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	kk, errRepo := controller.usecase.GetAllKemejaKeranjang(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no user column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(kk))
}

func (controller *KemejaKeranjangController) GetKemejaKeranjangDetail(c echo.Context) error {
	ctx := c.Request().Context()

	kkId, _ := strconv.Atoi(c.Param("kkId"))

	if kkId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja keranjang ID empty")
	}

	kk, errRepo := controller.usecase.GetKemejaKeranjangDetail(kkId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kk))
}

func (controller *KemejaKeranjangController) EditKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var kk request.KemejaKeranjangEdit
	kkId, _ := strconv.Atoi(c.Param("kkId"))
	err := c.Bind(&kk)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if kkId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja keranjang ID empty")
	}

	if kk.Jumlah == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Jumlah empty")
	}

	if kk.Size == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Size empty")
	}

	kkRepo, errRepo := controller.usecase.EditKemejaKeranjang(*kk.ToUsecase(), kkId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kkRepo))
}

func (controller *KemejaKeranjangController) DeleteKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	kkId, _ := strconv.Atoi(c.Param("kkId"))

	if kkId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja keranjang ID empty")
	}

	kk, errRepo := controller.usecase.DeleteKemejaKeranjang(kkId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kk))
}
