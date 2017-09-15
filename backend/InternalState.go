package main

type InternalState struct {
	CharStatValue CharStat
	CharVarValue  CharVar
	Projects      []Project
	ActiveEffects []ActiveEffect
}
