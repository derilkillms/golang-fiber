package migration

import (
	"fmt"
	"golang-fiber/database"
	"golang-fiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)

	}
	fmt.Println("Database Migrated")
}
