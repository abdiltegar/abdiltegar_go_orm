package main

import (
	"learn_orm/config"
	"learn_orm/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8083"))
}
