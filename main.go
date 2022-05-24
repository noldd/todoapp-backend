package main

import (
	"todoapp-backend/app"
	"todoapp-backend/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config.GetConfig()
	app := app.NewApp(config)
	app.Run(":" + config.Port)
}
