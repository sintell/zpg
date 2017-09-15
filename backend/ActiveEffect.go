package main

type ActiveEffect struct {
	CharStatId int      `json:"-"`
	CharStat   CharStat `json:"-"`
	Expires    int      `json:"expires"`
	EffectId   int      `json:"-"`
	Effect     *Effect
}
