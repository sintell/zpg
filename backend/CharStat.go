package main

type CharID int

type CharStat struct {
	ID      CharID `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	UserID  int    `json:"-"`
	User    *User  `json:"-"`
}
