package main

import "flag"

var configPath = flag.String("config", "config.json", "path to config.json")

func init() {
	initDB()
	runMigrations()
	initGlobalState()
}

func main() {
	GetGlobalState().save()
	server := GetServer()
	server.Run()
}
