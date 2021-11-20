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
	// users.Use(middleware.JWTWithConfig(controller.JWTConfig))
	e.POST("/login", controller.UserController.LoginController)
	e.POST("/signup", controller.UserController.SignUpController)
	// e.POST("/logout", users.LogoutContorller)
	eUser.PUT("/:userId", controller.UserController.EditUserController)
	eUser.DELETE("/:userId", controller.UserController.DeleteUserController)
	eUser.GET("/:userId", controller.UserController.GetDetailUserController)
	e.GET("/users", controller.UserController.GetAllUsersController)
	// users.GET("/", controller.UserController.Login, middleware.JWTWithConfig(controller.JWTConfig))
}
