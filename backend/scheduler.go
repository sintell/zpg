package main

import (
	"time"
)

type Scheduler struct {
}

func (ps *Scheduler) Start() {
	config := &SchedulerConfig{}
	ReadConfig(config, *configPath)
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
