package main

func init() {
	initDB()
	runMigrations()
	initGlobalState()
}

func main() {
	server := GetServer()
	server.Run()
}
