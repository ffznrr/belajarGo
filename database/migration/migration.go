package migration

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"log"
)

func RunMigration() {
	err:=database.DB.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated...")
}