package main

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
)

var db *pg.DB

func initDB() *pg.DB {
	dbOpts, err := pg.ParseURL("postgres://postgres:postgres@10.208.10.94/zpg")
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
