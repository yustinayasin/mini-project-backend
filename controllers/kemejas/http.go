package controllers

import (
	"kemejaku/business/kemejas"
	"kemejaku/controllers"
	"kemejaku/controllers/kemejas/request"
	"kemejaku/controllers/kemejas/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KemejaController struct {
	usecase kemejas.KemejaUseCaseInterface
}

func NewKemejaController(uc kemejas.KemejaUseCaseInterface) *KemejaController {
	return &KemejaController{
		usecase: uc,
	}
}

func (controller *KemejaController) InsertKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	var kemeja request.Kemeja

	err := c.Bind(&kemeja)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if kemeja.Nama == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja name empty")
	}

	if kemeja.Deskripsi == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja description empty")
	}

	if kemeja.Harga == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja price empty")
	}

	if kemeja.Stock_L == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja stock for size L empty")
	}

	if kemeja.Stock_M == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja stock for size M empty")
	}

	if kemeja.Stock_S == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja stock for size S empty")
	}

	kemejaRepo, errRepo := controller.usecase.InsertKemeja(*kemeja.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Failed to insert new kemeja", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemejaRepo))
}

func (controller *KemejaController) GetAllKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	kemeja, errRepo := controller.usecase.GetAllKemeja(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no kemeja column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(kemeja))
}

func (controller *KemejaController) GetKemejaDetail(c echo.Context) error {
	ctx := c.Request().Context()

	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))

	if kemejaId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja ID empty")
	}

	kemeja, errRepo := controller.usecase.GetKemejaDetail(kemejaId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemeja))
}

func (controller *KemejaController) EditKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	var kemeja request.Kemeja
	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))
	err := c.Bind(&kemeja)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if kemejaId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja ID empty")
	}

	if kemeja.Nama == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja name empty")
	}

	if kemeja.Deskripsi == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja description empty")
	}

	if kemeja.Harga == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja price empty")
	}

	kemejaRepo, errRepo := controller.usecase.EditKemeja(*kemeja.ToUsecase(), kemejaId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemejaRepo))
}

func (controller *KemejaController) DeleteKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))

	if kemejaId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja ID empty")
	}

	kemeja, errRepo := controller.usecase.DeleteKemeja(kemejaId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemeja))
}
