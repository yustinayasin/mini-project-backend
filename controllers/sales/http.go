package sales

import (
	"kemejaku/business/sales"
	"kemejaku/controllers"
	"kemejaku/controllers/sales/request"
	"kemejaku/controllers/sales/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SaleController struct {
	usecase sales.SaleUseCaseInterface
}

func NewSaleController(uc sales.SaleUseCaseInterface) *SaleController {
	return &SaleController{
		usecase: uc,
	}
}

func (controller *SaleController) InsertSale(c echo.Context) error {
	ctx := c.Request().Context()

	var sale request.Sale

	err := c.Bind(&sale)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if sale.Percent == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale percent empty")
	}

	if sale.MinimumPembelian == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale minimum pembelian empty")
	}

	if sale.StartDate == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale start date empty")
	}

	if sale.EndDate == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale end date empty")
	}

	saleRepo, errRepo := controller.usecase.InsertSale(*sale.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Failed to insert new sale", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(saleRepo))
}

func (controller *SaleController) GetAllSale(c echo.Context) error {
	ctx := c.Request().Context()

	sale, errRepo := controller.usecase.GetAllSale(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no sale column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(sale))
}

func (controller *SaleController) GetSaleDetail(c echo.Context) error {
	ctx := c.Request().Context()

	saleId, _ := strconv.Atoi(c.Param("saleId"))

	if saleId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale ID empty")
	}

	sale, errRepo := controller.usecase.GetSaleDetail(saleId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Sale not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(sale))
}

func (controller *SaleController) EditSale(c echo.Context) error {
	ctx := c.Request().Context()

	var sale request.Sale
	saleId, _ := strconv.Atoi(c.Param("kemejaId"))
	err := c.Bind(&sale)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	if saleId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale ID empty")
	}

	if sale.Percent == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale percent empty")
	}

	if sale.MinimumPembelian == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale minimum pembelian empty")
	}

	if sale.StartDate == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale start date empty")
	}

	if sale.EndDate == "" {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale end date empty")
	}

	saleRepo, errRepo := controller.usecase.EditSale(*sale.ToUsecase(), saleId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Sale not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(saleRepo))
}

func (controller *SaleController) DeleteSale(c echo.Context) error {
	ctx := c.Request().Context()

	saleId, _ := strconv.Atoi(c.Param("saleId"))

	if saleId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Sale ID empty")
	}

	sale, errRepo := controller.usecase.DeleteSale(saleId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Sale not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(sale))
}