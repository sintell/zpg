package main

type ActiveEffects struct {
	tableName     struct{} `sql:"active_effects"`
	CharStatId    int
	CharStatValue CharStat
	Expires       int
	EffectId      int
	Effect        *Project
}
