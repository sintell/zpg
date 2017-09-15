package main

type ActiveEffect struct {
	CharStatId    int      `json:"-"`
	CharStatValue CharStat `json:"-"`
	Expires       int      `json:"expires"`
	EffectId      int      `json:"-"`
	Effect        *Effect
}
