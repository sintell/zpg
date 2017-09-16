package main

type ActiveEffect struct {
	CharStatID CharID   `json:"-"`
	CharStat   CharStat `json:"-"`
	Expires    int      `json:"expires"`
	EffectID   int      `json:"-"`
	Effect     *Effect
}
