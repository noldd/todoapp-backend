package integration

import (
	"log"
	"math/rand"
	"time"
	"todoapp-backend/app/model"

	"gorm.io/gorm"
)

func randomTitle() string {
    charSet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    length := 16
    b := make([]rune, length)
    for i := range b {
        b[i] = charSet[rand.Intn(len(charSet))]
    }
    return string(b)
}

func randomBool() bool {
    return rand.Intn(2) == 1
}

// Creates a random task to DB and returns a reference to it.
func randomExistingTask(db *gorm.DB) (*model.Task) {
    task := &model.Task{
        Title: randomTitle(),
        Done: randomBool(),
    }

    if err := db.Save(task).Error; err != nil {
        log.Fatalf("Failed save task to DB: %v", err)
    }

    return task
}

// Creates a random task that doesn't exist in the DB and returns reference.
func randomNonExistingTask(db *gorm.DB) (*model.Task) {
    task := randomExistingTask(db)
    db.Delete(&task)
    return task
}

func init() {
    rand.Seed(time.Now().UnixNano())
}
