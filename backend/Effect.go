package main

type Effect struct {
	Id          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	EffectId    int `json:"-"`
	EffectValue *SkillValue	`json:"values"`
}
