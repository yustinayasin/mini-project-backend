package routes

import (
	kemejaKeranjangController "kemejaku/controllers/kemejakeranjangs"
	kemejaController "kemejaku/controllers/kemejas"
	keranjangController "kemejaku/controllers/keranjangs"
	userController "kemejaku/controllers/users"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UserController            userController.UserController
	KeranjangController       keranjangController.KeranjangController
	KemejaController          kemejaController.KemejaController
	KemejaKeranjangController kemejaKeranjangController.KemejaKeranjangController
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

	eKeranjang := e.Group("/keranjang")

	eKeranjang.GET("/:keranjangId", controller.KeranjangController.GetKeranjangDetail)
	e.GET("/keranjangs", controller.KeranjangController.GetAllKeranjang)
	e.POST("/keranjangs", controller.KeranjangController.InsertKeranjang)
	eKeranjang.PUT("/:keranjangId", controller.KeranjangController.EditKeranjang)
	eKeranjang.DELETE("/:keranjangId", controller.KeranjangController.DeleteKeranjang)

	eKemejaKeranjang := e.Group("/kemejakeranjang")

	eKemejaKeranjang.GET("/:kemejaKeranjangId", controller.KemejaKeranjangController.GetKemejaKeranjangDetail)
	e.GET("/kemejakeranjangs", controller.KemejaKeranjangController.GetAllKemejaKeranjang)
	e.POST("/kemejakeranjangs", controller.KemejaKeranjangController.InsertKemejaKeranjang)
	eKemejaKeranjang.PUT("/:kemejaKeranjangId", controller.KemejaKeranjangController.EditKemejaKeranjang)
	eKemejaKeranjang.DELETE("/:kemejaKeranjangId", controller.KemejaKeranjangController.DeleteKemejaKeranjang)

	eKemeja := e.Group("/kemeja")

	eKemeja.GET("/:kemejaKeranjangId", controller.KemejaController.GetKemejaDetail)
	e.GET("/kemejakeranjangs", controller.KemejaController.GetAllKemeja)
	e.POST("/kemejakeranjangs", controller.KemejaController.InsertKemeja)
	eKemeja.PUT("/:kemejaKeranjangId", controller.KemejaController.EditKemeja)
	eKemeja.DELETE("/:kemejaKeranjangId", controller.KemejaController.DeleteKemeja)
}
