package main

type SkillValues struct {
	tableName struct{} `sql:"skill_values"`
	Id int
	Prog int
	Testing int
	Analyze int
}
