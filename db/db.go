package db

import (
	"fmt"
	"log"
	"todoapp-backend/app/model"
	"todoapp-backend/config"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB(config *config.Config) *gorm.DB {
	dbURI := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Endpoint,
		config.DB.Name,
		config.DB.Charset,
	)

	dbDialector := mysql.New(mysql.Config{
		DSN: dbURI,
	})

	db, err := gorm.Open(dbDialector)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	model.DBMigrate(db)
    return db
}
