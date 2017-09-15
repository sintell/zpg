package main

type InternalState struct {
	CharStatValue CharStat
	CharVarValue  CharVar
	Projects      []Project
	ActiveEffects []ActiveEffect
}

func of(Id int) InternalState {
	result := InternalState{}
	GetDB().Model(&result.CharStatValue).Where("user_id = ?", Id).Select()
	charId := result.CharStatValue.Id
	GetDB().Model(&result.CharVarValue).Where("char_stat_id = ?", charId).Select()
	GetDB().Model(&result.Projects).Where("char_stat_id = ?", charId).Select()
	GetDB().Model(&result.ActiveEffects).Where("char_stat_id = ?", charId).Select()
	return result
}
