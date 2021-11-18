package configs

import (
	"kemejaku/models/kemejakeranjangs"
	"kemejaku/models/kemejas"
	"kemejaku/models/keranjangs"
	"kemejaku/models/sales"
	"kemejaku/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//dibuat global biar bisa diakses di dalam 2 function
var DB *gorm.DB

func ConnectDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/kemejaku?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database tidak connect")
	}
	Migration()
}

func Migration() {
	// auto migrate
	DB.AutoMigrate(kemejas.Kemeja{}, users.User{}, sales.Sale{}, keranjangs.Keranjang{}, kemejakeranjangs.KemejaKeranjang{})
}
