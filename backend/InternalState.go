package main

type InternalState struct {
	CharStatValue *CharStat
	CharVarValue  *CharVar
	Projects      []*Project
	ActiveEffects []*ActiveEffect
	EventQueue    *EventQueue
}

func StateFromDB(id CharID) *InternalState {
	cv := &CharVar{CharStatID: id}
	var p []*Project
	var e []*ActiveEffect

	GetDB().Model(cv).Column("char_var.*", "CharStat", "CurrentProject", "SkillValue").Select()
	GetDB().Model(cv).Where("char_stat_id = ?", id).Select()
	GetDB().Model(&p).Column("project.*", "ProgrValues", "ReqValues").Where("char_stat_id = ?", id).Select(&p)
	GetDB().Model(&e).Where("char_stat_id = ?", id).Select()

	state := &InternalState{
		CharStatValue: cv.CharStat,
		CharVarValue:  cv,
		Projects:      p,
		ActiveEffects: e,
		EventQueue:    NewEventQueue(),
	}
	return state
}

func NewInternalState(csv *CharStat, cvv *CharVar) *InternalState {
	prj, _ := CreateProjectsFor(cvv)
	prj[0].Active = true
	prj[0].Status = Analyze
	cvv.CurrentProjectID = prj[0].ID
	cvv.CurrentProject = prj[0]

	cvv.Save()
	return &InternalState{CharStatValue: csv, CharVarValue: cvv, Projects: prj, EventQueue: NewEventQueue()}
}
