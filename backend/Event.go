package main

import "time"

type Event struct {
	eventType   EventType
	description string
	timestamp   time.Time
}

type EventType string

const (
	New_project           EventType = "NEW_PROJECT"
	Change_project_status           = "CHANGE_PROJECT_STATUS"
	Complete_project                = "COMPLETE_PROJECT"
	New_event                       = "NEW_EVENT"
	New_effect                      = "NEW_EFFECT"
	Level_up                        = "LEVEL_UP"
	Effect_expires                  = "EFFECT_EXPIRES"
)

func NewEvent(eventType EventType, description string) *Event {
	return &Event{eventType: eventType, description: description, timestamp: time.Now()}
}
