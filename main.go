package main

import (
	"webserver/controllers"
	"webserver/utils"
)

func main() {
	utils.ClearScreen()

	controllers.StartMainServer()
}