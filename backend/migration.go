package main

import "github.com/go-pg/pg/orm"

func runMigrations() {
	db := GetDB()

	db.DropTable(&Project{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&SkillValue{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&ActiveEffect{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&CharStat{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&CharVar{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&Effect{}, &orm.DropTableOptions{IfExists: true})

	db.CreateTable(&User{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&CharStat{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&CharVar{}, &orm.CreateTableOptions{IfNotExists: true})
	db.CreateTable(&Project{}, nil)
	db.CreateTable(&SkillValue{}, nil)
	db.CreateTable(&ActiveEffect{}, nil)
	db.CreateTable(&Effect{}, nil)
}
