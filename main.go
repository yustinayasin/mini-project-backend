package main

import (
	"kemejaku/configs"
	"kemejaku/routes"
)

func main() {
	configs.ConnectDb()
	e := routes.New()
	e.Start(":8000")
}
