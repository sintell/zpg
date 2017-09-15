package main

func init() {
	runMigrations()
}

func main() {
	server := GetServer()
	server.Run()
}
