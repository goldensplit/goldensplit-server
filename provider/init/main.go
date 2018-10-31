// Copyright 2018 Kirill Lapshin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/goldensplit/gs-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	createDB()

	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=goldensplit sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dropTables(db)
	createTables(db)
}

func createDB() {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("CREATE DATABASE goldensplit;")
}

func createTables(db *gorm.DB) error {
	tx := db.Begin()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.CreateTable(&model.User{}, &model.Session{}, &model.Video{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func dropTables(db *gorm.DB) error {
	tx := db.Begin()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.DropTable(&model.User{}, &model.Session{}, &model.Video{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
