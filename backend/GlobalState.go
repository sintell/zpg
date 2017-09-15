package main

var globalState *GlobalState

type GlobalState struct {
	states map[int]InternalState
}

func (state GlobalState) get(Id int) InternalState {
	return state.states[Id]
}

func (state GlobalState) getIds() []int {
	result := make([]int, len(state.states))
	counter := 0
	for id := range state.states {
		result[counter] = id
	}
	return result
}

func (state GlobalState) save() {
	for _, value := range state.states {
		GetDB().Model(value.CharStatValue).OnConflict("(id) DO UPDATE").Insert()
		GetDB().Model(value.CharVarValue).OnConflict("(id) DO UPDATE").Insert()
		GetDB().Model(value.Projects).OnConflict("(id) DO UPDATE").Insert()
		GetDB().Model(value.ActiveEffects).OnConflict("(id) DO UPDATE").Insert()
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
	globalState = &GlobalState{make(map[int]InternalState)}
	globalState.load()
	return globalState
}
