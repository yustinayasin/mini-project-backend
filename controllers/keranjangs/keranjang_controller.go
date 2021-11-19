package keranjangs

import (
	"kemejaku/configs"
	"kemejaku/models/keranjangs"
	"kemejaku/models/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertKeranjangController(c echo.Context) error {
	var keranjang keranjangs.Keranjang
	c.Bind(&keranjang)

	if keranjang.UserId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id Kosong",
			nil,
		})
	}

	keranjang.Status = false

	//orm akan otomatis ngecek berdasarkan nama tabel dan dimasukkan ke variable news
	result := configs.DB.Create(&keranjang)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		keranjang,
	}

	return c.JSON(http.StatusOK, response)
}

func GetDetailKeranjangController(c echo.Context) error {
	//mengambil parameter dengan key yang sama
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	if keranjangId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id Kosong",
			nil,
		})
	}

	var data = keranjangs.Keranjang{}

	configs.DB.Preload("KemejaKeranjang").First(&data, keranjangId)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllKeranjangController(c echo.Context) error {
	var keranjangs []keranjangs.Keranjang

	result := configs.DB.Preload("KemejaKeranjang").Find(&keranjangs)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		keranjangs,
	}

	return c.JSON(http.StatusOK, response)
}

func EditKeranjangController(c echo.Context) error {
	var keranjang keranjangs.Keranjang
	var newKeranjang keranjangs.Keranjang
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))
	c.Bind(&keranjang)

	if keranjangId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id kosong",
			nil,
		})
	}

	configs.DB.Preload("KemejaKeranjang").First(&newKeranjang, keranjangId)

	newKeranjang.Status = keranjang.Status

	configs.DB.Save(&newKeranjang)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		newKeranjang,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteKeranjangController(c echo.Context) error {
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	var data = keranjangs.Keranjang{}

	if keranjangId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id Kosong",
			nil,
		})
	}

	configs.DB.Delete(&data, keranjangId)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		nil,
	}
	return c.JSON(http.StatusOK, response)
}
