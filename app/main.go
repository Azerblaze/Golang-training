package main

import (
	"training/app/configs"
	"training/app/database"
)

func main() {
	configs.InitConfig()
	database.InitDatabase()
}
