package users

import (
	"kemejaku/configs"
	"kemejaku/middlewares"
	"kemejaku/models/responses"
	"kemejaku/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	// email := c.FormValue("email")
	// password := c.FormValue("password")

	//binding data
	var user users.User
	c.Bind(&user)

	if err := configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, responses.BaseResponse{
			http.StatusUnauthorized,
			"Email dan Password tidak sesuai",
			nil,
		})
	}

	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Email Kosong",
			nil,
		})
	}

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	//generate token
	token, err := middlewares.GenerateTokenJWT(int(user.Id))

	//sukses login
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			http.StatusInternalServerError,
			"Error ketika generate token JWT",
			nil,
		})
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		http.StatusOK,
		"success",
		map[string]string{
			"token": token,
		},
	})
}

func SignUpController(c echo.Context) error {
	var user users.User
	c.Bind(&user)

	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Email Kosong",
			nil,
		})
	}

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	//orm akan otomatis ngecek berdasarkan nama tabel dan dimasukkan ke variable news
	result := configs.DB.Create(&user)

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
		user,
	}

	return c.JSON(http.StatusOK, response)
}

// func EditUserController(c echo.Context) error {
// 	var user user.User

// }

func GetDetailUserController(c echo.Context) error {
	//mengambil parameter dengan key yang sama
	userId, _ := strconv.Atoi(c.Param("userId"))

	if userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"User Id Kosong",
			nil,
		})
	}

	var data = users.User{}

	configs.DB.First(&data, userId)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllUserController(c echo.Context) error {
	var users []users.User

	result := configs.DB.Find(&users)

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
		users,
	}

	return c.JSON(http.StatusOK, response)
}

func EditUserController(c echo.Context) error {
	var user users.User
	var newUser users.User
	userId, _ := strconv.Atoi(c.Param("userId"))
	c.Bind(&user)

	if userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"User Id kosong",
			nil,
		})
	}

	configs.DB.First(&newUser, userId)

	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Email Kosong",
			nil,
		})
	} else {
		newUser.Email = user.Email
	}

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	} else {
		newUser.Password = user.Password
	}

	if user.Name != "" {
		newUser.Name = user.Name
	}

	if user.PhoneNumber != "" {
		newUser.PhoneNumber = user.PhoneNumber
	}

	if user.Street != "" {
		newUser.Street = user.Street
	}

	if user.Address != "" {
		newUser.Address = user.Address
	}

	if user.PostalCode != "" {
		newUser.PostalCode = user.PostalCode
	}

	configs.DB.Save(&newUser)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		newUser,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))

	var data = users.User{}

	if userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			http.StatusBadRequest,
			"User Id Kosong",
			nil,
		})
	}

	configs.DB.Delete(&data, userId)

	var response = responses.BaseResponse{
		http.StatusOK,
		"success",
		nil,
	}
	return c.JSON(http.StatusOK, response)
}
