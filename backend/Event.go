package main

import "time"

type Event struct {
	name        string
	eventType   string
	description string
	timestamp   time.Time
}
