package main

import "github.com/go-pg/pg/orm"

func runMigrations() {
	db := GetDB()

	db.DropTable(&User{}, &orm.DropTableOptions{IfExists: true})
	db.CreateTable(&User{}, nil)
}
