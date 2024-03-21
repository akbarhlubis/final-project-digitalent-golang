package main

import (
	"final-project-akbar/config"
	"final-project-akbar/router"
)

func main() {
	r := router.SetupRouter()
	// Menjalankan server pada port 8080
	port := config.AppPort
	r.Run(":" + port)
}
