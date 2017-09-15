package main

type InternalState struct {
	CharStatValue *CharStat
	CharVarValue  *CharVar
	Projects      []*Project
	ActiveEffects []*ActiveEffect
}

func StateFromDB(id CharID) *InternalState {
	state := &InternalState{}
	GetDB().Model(&state.CharVarValue).Where("char_stat_id = ?", id).Select()
	GetDB().Model(&state.Projects).Where("char_stat_id = ?", id).Select()
	GetDB().Model(&state.ActiveEffects).Where("char_stat_id = ?", id).Select()
	return state
}

func NewInternalState(csv *CharStat, cvv *CharVar) *InternalState {
	return &InternalState{CharStatValue: csv, CharVarValue: cvv, Projects: CreateProjectsFor(csv.ID)}
}
