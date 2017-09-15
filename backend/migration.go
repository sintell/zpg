package main

import "github.com/go-pg/pg/orm"

func runMigrations() {
	db := GetDB()

	db.DropTable(&User{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&CharStat{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&Projects{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&SkillValues{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&CharVar{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&ActiveEffects{}, &orm.DropTableOptions{IfExists: true})
	db.DropTable(&Effects{}, &orm.DropTableOptions{IfExists: true})

	db.CreateTable(&User{}, nil)
	db.CreateTable(&CharStat{}, nil)
	db.CreateTable(&Projects{}, nil)
	db.CreateTable(&SkillValues{}, nil)
	db.CreateTable(&CharVar{}, nil)
	db.CreateTable(&ActiveEffects{}, nil)
	db.CreateTable(&Effects{}, nil)
}
