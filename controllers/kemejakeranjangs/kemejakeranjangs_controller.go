package kemejakeranjangs

import (
	"kemejaku/configs"
	"kemejaku/models/kemejakeranjangs"
	"kemejaku/models/responses"
	"kemejaku/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertKemejaKeranjangController(c echo.Context) error {
	var kemejakeranjang kemejakeranjangs.KemejaKeranjang
	c.Bind(&kemejakeranjang)

	if kemejakeranjang.IdKemeja == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Kemeja Id Kosong",
			nil,
		})
	}

	if kemejakeranjang.IdKeranjang == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id Kosong",
			nil,
		})
	}

	if kemejakeranjang.Jumlah == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Jumlah Kosong",
			nil,
		})
	}

	if kemejakeranjang.Size == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Size Kosong",
			nil,
		})
	}

	result := configs.DB.Create(&kemejakeranjang)

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
		kemejakeranjang,
	}

	return c.JSON(http.StatusOK, response)
}

func GetDetailKemejaKeranjangController(c echo.Context) error {
	//mengambil parameter dengan key yang sama
	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	if kemejaId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Kemeja Id Kosong",
			nil,
		})
	}

	if keranjangId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id Kosong",
			nil,
		})
	}

	var data = users.User{}

	configs.DB.First(&data, kemejaId, keranjangId)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllKemejaKeranjangController(c echo.Context) error {
	var kemejakeranjangs []kemejakeranjangs.KemejaKeranjang

	result := configs.DB.Find(&kemejakeranjangs)

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
		kemejakeranjangs,
	}

	return c.JSON(http.StatusOK, response)
}

func EditKemejaKeranjangController(c echo.Context) error {
	var kemejakeranjang kemejakeranjangs.KemejaKeranjang
	var newKemejakeranjang kemejakeranjangs.KemejaKeranjang
	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))
	c.Bind(&kemejakeranjang)

	if kemejaId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Kemeja Id Kosong",
			nil,
		})
	}

	if keranjangId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id Kosong",
			nil,
		})
	}

	configs.DB.First(&newKemejakeranjang, kemejaId, keranjangId)

	if kemejakeranjang.Jumlah == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Jumlah Kosong",
			nil,
		})
	} else {
		newKemejakeranjang.Jumlah = kemejakeranjang.Jumlah
	}

	configs.DB.Save(&newKemejakeranjang)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		newKemejakeranjang,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteKemejaKeranjangController(c echo.Context) error {
	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	var kemejakeranjang = kemejakeranjangs.KemejaKeranjang{}

	if kemejaId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Kemeja Id Kosong",
			nil,
		})
	}

	if keranjangId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Keranjang Id Kosong",
			nil,
		})
	}

	configs.DB.Delete(&kemejakeranjang, kemejaId, keranjangId)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		nil,
	}
	return c.JSON(http.StatusOK, response)
}
