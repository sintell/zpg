package main

type ActiveEffect struct {
	CharStatID int      `json:"-"`
	CharStat   CharStat `json:"-"`
	Expires    int      `json:"expires"`
	EffectID   int      `json:"-"`
	Effect     *Effect
}
