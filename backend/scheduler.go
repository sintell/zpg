package main

import (
	"time"
)

type persistSheduler struct {
}

func (ps *persistSheduler) Start() {
	config := &SchedulerConfig{}
	ReadConfig(config, "backend/resources/config.json")
	persistTicker := time.NewTicker(time.Millisecond * config.PersistInterval)
	go func() {
		for _ := range persistTicker.C {
			GetGlobalState().save()
		}
	}()
}
