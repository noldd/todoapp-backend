package main

import (
	"todoapp-backend/app"
	"todoapp-backend/config"
)

func main() {
	config := config.GetConfig()
	app := app.NewApp(config)
	app.Run(":" + config.Port)
}
