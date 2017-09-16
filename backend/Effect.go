package main

type Effect struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"desc"`
	ValueID     int         `json:"-"`
	Value       *SkillValue `json:"values"`
}
