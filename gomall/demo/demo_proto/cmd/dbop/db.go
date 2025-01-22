package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/readlnh/biz-demo/gomall/demo/demo-proto/biz/dal"
	"github.com/readlnh/biz-demo/gomall/demo/demo-proto/biz/dal/mysql"
	"github.com/readlnh/biz-demo/gomall/demo/demo-proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CURD
	// Created
	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "123456"})

	// Update
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "654321")

	// Read
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)

	fmt.Printf("row: %+v\n", row)

	// Delete
	mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})
}
