package main

import "time"

type EventQueueElement struct {
	EventType   EventType `json:"event_type"`
	Description string    `json:"desc"`
	Timestamp   time.Time `json:"timestamp"`
	Order       int       `json:"order"`
}

type ExternalState struct {
	Character struct {
		*CharStat
		*CharVar
	} `json:"char"`
	Projects      []*Project           `json:"projects"`
	ActiveEffects []*ActiveEffect      `json:"effects"`
	Events        []*EventQueueElement `json:"events"`
}

func from(isv *InternalState) *ExternalState {
	result := &ExternalState{}
	result.Character.CharStat = isv.CharStatValue
	result.Character.CharVar = isv.CharVarValue
	result.Projects = isv.Projects
	result.ActiveEffects = isv.ActiveEffects
	for _, effect := range result.ActiveEffects {
		result.Character.CharVar.SkillValue.Prog += effect.Effect.Effect.Prog
		result.Character.CharVar.SkillValue.Testing += effect.Effect.Effect.Testing
		result.Character.CharVar.SkillValue.Analyze += effect.Effect.Effect.Analyze
	}
	result.Events = isv.EventQueue.toExternalForm()
	return result
}
