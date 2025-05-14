package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

	type User struct {
		ID        uint   `json:"id" gorm:"primaryKey" `
		Name      string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Address   string `json:"Address" validate:"required"`
		Phone     string `json:"Phone"`
		Email     string `json:"Email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	type Product struct {
		ID        	uint   `json:"id" gorm:"primaryKey" `
		Name      	string `json:"name" validate:"required"`
		Price 		int `json:"price"`
		Stock 		int `json:"stock"`
		CreatedAt 	time.Time `json:"created_at"`
		UpdatedAt 	time.Time `json:"updated_at"`
		
	}


func (u User) GetUser(){
	fmt.Println("user")
}