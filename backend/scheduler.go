package main

import (
	"fmt"
	"time"
)

type Scheduler struct {
}

func (ps *Scheduler) Start() {
	config := &SchedulerConfig{}
	ReadConfig(config, *configPath)
	persistTicker := time.NewTicker(time.Millisecond * config.PersistInterval)
	go func() {
		for t := range persistTicker.C {
			GetGlobalState().save()
		}
	}()

	mainloopTicker := time.NewTicker(time.Millisecond * config.CalculateInterval)
	go func() {
		for t := range mainloopTicker.C {
			MainloopTick()
		}
	}()
}
