package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)
	defer fmt.Println("Successfully Migrated")
	if err := dbConn.AutoMigrate(&model.Article{}, &model.Comment{}, &model.Reply{}); err != nil {
		fmt.Println("Error Migrating")
	}
}
