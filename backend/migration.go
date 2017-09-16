package main

import "github.com/go-pg/pg/orm"

func runMigrations() {
	db := GetDB()

	db.CreateTable(&User{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&CharStat{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&CharVar{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&Project{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&SkillValue{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&ActiveEffect{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&Effect{}, &orm.CreateTableOptions{IfNotExists: true})
}
