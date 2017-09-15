package main

type CharStat struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	UserId  int    `json:"-"`
	User    *User  `json:"-"`
}
