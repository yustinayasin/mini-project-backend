package routes

import (
	userController "kemejaku/controllers/users"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UserController userController.UserController
	// JWTConfig      middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {
	eUser := e.Group("/user")
	// eUser.Use(middleware.JWTWithConfig(controller.JWTConfig))
	e.POST("/login", controller.UserController.Login)
	e.POST("/signup", controller.UserController.SignUp)
	// e.POST("/logout", users.LogoutContorller)
	eUser.PUT("/:userId", controller.UserController.EditUser)
	eUser.DELETE("/:userId", controller.UserController.DeleteUser)
	eUser.GET("/:userId", controller.UserController.GetUserDetail)
	e.GET("/users", controller.UserController.GetAllUsers)
	// users.GET("/", controller.UserController.Login, middleware.JWTWithConfig(controller.JWTConfig))
}
