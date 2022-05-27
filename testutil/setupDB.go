package testutil

import (
	"log"
	"todoapp-backend/app/model"
	"todoapp-backend/config"
	"todoapp-backend/db"
)

// Fills DB with data
func SetupDB() {
    SetEnv()
    db := db.GetDB(config.GetConfig())

    // Tasks
    tasks := []model.Task{
        { Title: "Foo", Done: true },
        { Title: "Bar", Done: false },
    }

    for _, task := range tasks {
        if err := db.Save(&task).Error; err != nil {
            log.Fatalf("Failed to save task: %v", err)
        }
    }
}
