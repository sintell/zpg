package main

var globalState *GlobalState

type GlobalState struct {
	states map[int]InternalState
}

func (state GlobalState) save() {
	for _, value := range state.states {
		GetDB().Model(value.CharStatValue).OnConflict("(id) DO UPDATE").Insert(value.CharStatValue)
		GetDB().Model(value.CharVarValue).OnConflict("(id) DO UPDATE").Insert(value.CharStatValue)
		GetDB().Model(value.Projects...).OnConflict("(id) DO UPDATE").Insert(value.CharStatValue)
		GetDB().Model(value.ActiveEffects...).OnConflict("(id) DO UPDATE").Insert(value.CharStatValue)
	}
}

func (state GlobalState) load() GlobalState {
	Ids := []int{}
	GetDB().Model(&User{}).Column("id").Select(&Ids)
	for _, id := range Ids {
		state.states[id] = of(id)
	}
	return state
}

func GetGlobalState() *GlobalState {
	if globalState != nil {
		return globalState
	}

	return initGlobalState()
}

func initGlobalState() *GlobalState {
	globalState = &GlobalState{}
	globalState.load()
	return globalState
}
