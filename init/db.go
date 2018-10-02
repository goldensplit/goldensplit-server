package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	createDB()
}

func createDB() {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("CREATE DATABASE goldensplit;")
}