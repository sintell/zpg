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
	Level_up                        = "LEVEL_UP"
)

func NewEvent(eventType EventType, description string) *Event {
	return &Event{eventType: eventType, description: description, timestamp: time.Now()}
}
