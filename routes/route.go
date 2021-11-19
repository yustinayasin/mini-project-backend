package routes

import (
	"kemejaku/constants"
	"kemejaku/controllers/kemejakeranjangs"
	"kemejaku/controllers/keranjangs"
	"kemejaku/controllers/users"
	"kemejaku/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// kenapa return pointer
func New() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	middlewares.LogMiddleware(e)

	eUser := e.Group("/user")

	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte(constants.SECRET_JWT),
	}
	eUser.Use(middleware.JWTWithConfig(config))

	//routing
	e.POST("/login", users.LoginController)
	e.POST("/signup", users.SignUpController)
	// e.POST("/logout", users.LogoutContorller)
	eUser.PUT("/:userId", users.EditUserController)
	eUser.DELETE("/:userId", users.DeleteUserController)
	eUser.GET("/:userId", users.GetDetailUserController)
	e.GET("/users", users.GetAllUserController)

	// e.GET("/kemeja:kemejaId", kemeja.GetDetailkemejaController)
	// e.GET("/kemejas", kemeja.GetAllKemejaController)
	// e.POST("/kemejas", kemeja.InsertKemejaController)
	// e.PUT("/kemeja:kemejaId", kemeja.EditKemejaController)
	// e.DELETE("/kemeja:kemejaId", kemeja.DeleteKemejaController)

	eKeranjang := e.Group("/keranjang")

	eKeranjang.Use(middleware.JWTWithConfig(config))

	eKeranjang.GET("/:keranjangId", keranjangs.GetDetailKeranjangController)
	e.GET("/keranjangs", keranjangs.GetAllKeranjangController)
	e.POST("/keranjangs", keranjangs.InsertKeranjangController)
	eKeranjang.PUT("/:keranjangId", keranjangs.EditKeranjangController)
	eKeranjang.DELETE("/:keranjangId", keranjangs.DeleteKeranjangController)

	eKemejaKeranjang := e.Group("/kemejakeranjang")

	eKemejaKeranjang.Use(middleware.JWTWithConfig(config))

	eKemejaKeranjang.GET("/:kemejaKeranjangId", kemejakeranjangs.GetDetailKemejaKeranjangController)
	e.GET("/kemejakeranjangs", kemejakeranjangs.GetAllKemejaKeranjangController)
	e.POST("/kemejakeranjangs", kemejakeranjangs.InsertKemejaKeranjangController)
	eKemejaKeranjang.PUT("/:kemejaKeranjangId", kemejakeranjangs.EditKemejaKeranjangController)
	eKemejaKeranjang.DELETE("/:kemejaKeranjangId", kemejakeranjangs.DeleteKemejaKeranjangController)

	return e
}
