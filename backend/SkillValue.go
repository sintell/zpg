package main

type SkillValue struct {
	Id      int `json:"-"`
	Prog    int `json:"prog"`
	Testing int `json:"test"`
	Analyze int `json:"analyze"`
}
