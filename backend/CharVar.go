package main

type CharVar struct {
	CharStatId       int
	CharStatValue    CharStat
	Stress           int
	Resting          int
	CurrentProjectId int
	CurrentProject   *Project
	SkillValueId     int
	SkillValue       *SkillValue
}
