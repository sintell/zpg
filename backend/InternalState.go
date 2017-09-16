package main

type InternalState struct {
	CharStatValue *CharStat
	CharVarValue  *CharVar
	Projects      []*Project
	ActiveEffects []*ActiveEffect
}

func StateFromDB(id CharID) *InternalState {
	cv := &CharVar{CharStatID: id}
	var p []*Project
	var e []*ActiveEffect

	GetDB().Model(cv).Column("char_val.*", "CharStat", "CurrentProject", "SkillValue").Select()
	GetDB().Model(cv).Where("char_stat_id = ?", id).Select()
	GetDB().Model(p).Column("projects.*", "ReqValues", "ProgrValues").Where("char_stat_id = ?", id).Select()
	GetDB().Model(e).Where("char_stat_id = ?", id).Select()

	state := &InternalState{
		CharStatValue: cv.CharStat,
		CharVarValue:  cv,
		Projects:      p,
		ActiveEffects: e,
	}
	return state
}

func NewInternalState(csv *CharStat, cvv *CharVar) *InternalState {
	prj, _ := CreateProjectsFor(csv.ID)
	return &InternalState{CharStatValue: csv, CharVarValue: cvv, Projects: prj}
}
