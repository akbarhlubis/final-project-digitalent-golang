package main

import (
	"final-project-akbar/config"
	"final-project-akbar/router"
)

func main() {
	r := router.SetupRouter()
	port := config.AppPort
	config.DBInit()
	config.SyncDB()
	r.Run(":" + port)
}
