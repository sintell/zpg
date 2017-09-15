package main

type User struct {
	tableName struct{} `sql:"user"`
	Id int
	Login    string
	Password string
}
