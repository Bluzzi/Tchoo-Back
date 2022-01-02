package main

import (
	"MetaFriend/config"
	"MetaFriend/database"
	"MetaFriend/routes"
)

func main() {
	config.Load()
	database.Load()
	routes.Load()
}
