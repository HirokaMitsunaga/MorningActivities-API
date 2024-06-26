package main

import (
	"fmt"
	"go-api/db"
	"go-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{}, &model.Timeline{}, &model.Like{}, &model.Comment{})
}
