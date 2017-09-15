package main

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
)

var db *pg.DB

func initDB() *pg.DB {
	dbConf := &DBConfig{}
	ReadConfig(dbConf, "backend/resources/config.json")
	dbOpts, err := pg.ParseURL(dbConf.PgAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	pg.SetLogger(log.New(os.Stdout, "db: ", log.Ldate|log.Ltime|log.LUTC))

	db = pg.Connect(dbOpts)

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return db
}

func GetDB() *pg.DB {
	if db != nil {
		return db
	}

	return initDB()
}
