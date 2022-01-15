package main

import (
	"MetaFriend/config"
	"MetaFriend/database"
	"MetaFriend/lottery"
	"MetaFriend/routes"
	"MetaFriend/routine"
)

func main() {
	config.Load()
	lottery.LoadLotteryPrizes()
	database.Load()
	routine.StartRoutine()
	routes.Load()
}
