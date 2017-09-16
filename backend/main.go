package main

import "flag"

var configPath = flag.String("config", "config.json", "path to config.json")

func init() {
	flag.Parse()
	initDB()
	runMigrations()
	initGlobalState()
}

func main() {
	server := GetServer()
	(&Scheduler{}).Start()
	server.Run()
}
