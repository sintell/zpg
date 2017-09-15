package main

type Effect struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"desc"`
	EffectID    int         `json:"-"`
	Effect      *SkillValue `json:"values"`
}
