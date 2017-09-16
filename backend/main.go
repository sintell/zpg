package main

import (
	"flag"
	"fmt"
)

var configPath = flag.String("config", "config.json", "path to config.json")
var resourcesPath = flag.String("resources", "", "path to resource")

func init() {
	flag.Parse()
	initDB()
	runMigrations()
	initGlobalState()
	(&Scheduler{}).Start()
	initStaticValues()
	err := LoadProjectNamesFrom(fmt.Sprintf("%s/project_names.json", *resourcesPath))
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	server := GetServer()
	server.Run()
}

func initStaticValues() {
	effects := make([]Effect, 0)
	effects = append(effects, Effect{ID: 0, Name: "Effect name 1", Description: "Все характеристики повышены на 5 на 4 часа",
		Value: &SkillValue{Prog: 5, Analyze: 5, Testing: 5}})
	effects = append(effects, Effect{ID: 1, Name: "Effect name 2", Description: "Все характеристики понижены на 2 на 5 часов",
		Value: &SkillValue{Prog: -2, Analyze: -2, Testing: -2}})
	GetDB().Model(effects).Insert(effects)
}
