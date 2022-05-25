package model

import "gorm.io/gorm"

func DBMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Task{})
	if err != nil {
		return err
	}
	return nil
}
