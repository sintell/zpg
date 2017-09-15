package main

import "sync"

var globalState *GlobalState

type GlobalState struct {
	states map[CharID]*InternalState
	mx     *sync.RWMutex
}

func (gs GlobalState) get(id CharID) *InternalState {
	gs.mx.RLock()
	defer gs.mx.RUnlock()

	return gs.states[id]
}

func (gs GlobalState) getIds() []CharID {
	gs.mx.RLock()
	defer gs.mx.RUnlock()

	result := make([]CharID, len(gs.states))
	counter := 0
	for id := range gs.states {
		result[counter] = id
	}
	return result
}

func (gs GlobalState) add(id CharID, state *InternalState) {
	gs.mx.Lock()
	defer gs.mx.Unlock()

	gs.states[id] = state
}

func (gs GlobalState) save() {
	gs.mx.Lock()
	defer gs.mx.Unlock()

	for _, value := range gs.states {
		GetDB().Model(value.CharStatValue).OnConflict("(id) DO UPDATE").Insert()
		GetDB().Model(value.CharVarValue).OnConflict("(id) DO UPDATE").Insert()
		GetDB().Model(value.Projects).OnConflict("(id) DO UPDATE").Insert()
		GetDB().Model(value.ActiveEffects).OnConflict("(id) DO UPDATE").Insert()
	}
}

func (gs GlobalState) load() GlobalState {
	gs.mx.Lock()
	defer gs.mx.Unlock()

	Ids := []CharID{}
	GetDB().Model(&CharStat{}).Column("id").Select(&Ids)
	for _, id := range Ids {
		gs.states[id] = StateFromDB(id)
	}
	return gs
}

func GetGlobalState() *GlobalState {
	if globalState != nil {
		return globalState
	}

	return initGlobalState()
}

func initGlobalState() *GlobalState {
	globalState = &GlobalState{states: make(map[CharID]*InternalState), mx: &sync.RWMutex{}}
	globalState.load()
	return globalState
}
