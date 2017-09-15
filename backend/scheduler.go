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
		for range persistTicker.C {
			GetGlobalState().save()
		}
	}()

	mainloopTicker := time.NewTicker(time.Millisecond * config.CalculateInterval)
	go func() {
		for range mainloopTicker.C {
			MainloopTick()
		}
	}()
}
