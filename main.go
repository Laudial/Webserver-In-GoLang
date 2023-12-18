package main

import (
	"SensiCortex/controllers"
	"SensiCortex/utils"
)

func main() {
	utils.ClearScreen()

	controllers.StartMainServer()
}