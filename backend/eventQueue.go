package main

type EventQueue struct {
	internalArray []Event
	logDepth      int
}

func NewEventQueue() *EventQueue {
	config := &EventQueueConfig{}
	ReadConfig(config, *configPath)
	return &EventQueue{internalArray: make([]Event, 0, config.LogDepth), logDepth: config.LogDepth}
}

func (eq *EventQueue) push(event *Event) {

	for i := len(eq.internalArray) - 2; i > 0; i-- {
		eq.internalArray[i] = eq.internalArray[i-1]
	}
	eq.internalArray[0] = *event
}

func (eq *EventQueue) toExternalForm() []*EventQueueElement {
	result := make([]*EventQueueElement, 0, len(eq.internalArray))
	for index, value := range eq.internalArray {
		result = append(result, &EventQueueElement{name:value.name, description:value.description,
		timestamp: value.timestamp, order:index})
	}
	return result
}
