package main

import (
	"kemejaku/app/routes"
	userUsecase "kemejaku/business/users"
	userController "kemejaku/controllers/users"
	"kemejaku/drivers/databases/mysql"
	userRepo "kemejaku/drivers/databases/users"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	//viper digunakan untuk membaca file config.json
	viper.SetConfigFile("app/config/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service Run and Debug mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{})
}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDb.InitialDB()
	dbMigrate(db)

	// jwt := _middleware.ConfigJWT{
	// 	SecretJWT:       viper.GetString(`jwt.secret`),
	// 	ExpiresDuration: viper.GetInt(`jwt.expired`),
	// }

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface, timeoutContext)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController: *userControllerInterface,
		// JWTConfig:      jwt.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
