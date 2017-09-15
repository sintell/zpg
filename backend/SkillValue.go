package main

type SkillValue struct {
	ID      int `json:"-"`
	Prog    int `json:"prog"`
	Testing int `json:"test"`
	Analyze int `json:"analyze"`
}
