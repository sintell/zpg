package main

import "flag"

var configPath = flag.String("config", "config.json", "path to config.json")

func init() {
	flag.Parse()
	initDB()
	runMigrations()
	initGlobalState()
	(&Scheduler{}).Start()
	initStaticValues()
}

func main() {
	server := GetServer()
	server.Run()
}

func initStaticValues() {
	effects := make([]Effect, 0)
	effects = append(effects, Effect{ID:0, Name: "Effect name 1", Description: "Все характеристики повышены на 5 на 4 часа",
		Effect: &SkillValue{Prog: 5, Analyze: 5, Testing: 5}})
	effects = append(effects, Effect{ID:0, Name: "Effect name 2", Description: "Все характеристики понижены на 10 на 7 часов",
		Effect: &SkillValue{Prog: -10, Analyze: -10, Testing: -10}})
	GetDB().Model(effects).Insert(effects)
}
