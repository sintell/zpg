package main

import (
	"math"
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

	eq.internalArray = append(
		[]*Event{event},
		eq.internalArray[:int(math.Max(float64(eq.logDepth-1), float64(len(eq.internalArray))))]...,
	)
	eq.internalArray[0] = event
}

func (eq *EventQueue) toExternalForm() []*EventQueueElement {
	result := make([]*EventQueueElement, 0, len(eq.internalArray))
	for index, value := range eq.internalArray {
		result = append(result, &EventQueueElement{name: value.name, description: value.description,
			timestamp: value.timestamp, order: index})
	}
	return result
}
