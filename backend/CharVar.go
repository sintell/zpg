package main

type CharVar struct {
	CharStatId       int         `json:"-"`
	CharStatValue    CharStat    `json:"-"`
	Stress           int         `json:"stress"`
	Resting          int         `json:"resting"`
	CurrentProjectId int         `json:"-"`
	CurrentProject   *Project    `json:"-"`
	SkillValueId     int         `json:"-"`
	SkillValue       *SkillValue `json:"skills"`
	Level            int         `json:"level"`
	Experience       int         `json:"exp"`
}
