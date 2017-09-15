package main

type ActiveEffect struct {
	CharStatId    int
	CharStatValue CharStat
	Expires       int
	EffectId      int
	Effect        *Project
}
