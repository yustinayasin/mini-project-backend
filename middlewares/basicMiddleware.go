package middlewares

import (
	"kemejaku/configs"
	"kemejaku/models/users"

	"github.com/labstack/echo/v4"
)

func BasicAuth(email, password string, c echo.Context) (bool, error) {
	var db = configs.DB
	var user users.User

	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, nil
	}

	return true, nil
}
