package main

import (
	"sync"
)

type EventQueue struct {
	internalArray []*Event
	logDepth      int
	mx            *sync.RWMutex
}

func NewEventQueue() *EventQueue {
	config := &EventQueueConfig{}
	ReadConfig(config, *configPath)
	return &EventQueue{
		internalArray: make([]*Event, 0, config.LogDepth),
		logDepth:      config.LogDepth,
		mx:            &sync.RWMutex{},
	}
}

func (eq *EventQueue) push(event *Event) {
	eq.mx.Lock()
	defer eq.mx.Unlock()

	eq.internalArray = append(
		[]*Event{event},
		eq.internalArray[:]...,
	)
	eq.internalArray[0] = event
}

func (eq *EventQueue) toExternalForm() []*EventQueueElement {
	eq.mx.Lock()
	defer eq.mx.Unlock()

	result := make([]*EventQueueElement, 0, len(eq.internalArray))
	for index, value := range eq.internalArray {
		result = append(result, &EventQueueElement{EventType: value.eventType, Description: value.description,
			Timestamp: value.timestamp, Order: index})
	}
	return result
}
