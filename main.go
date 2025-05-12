package main

import (
	"good_wave_back_end/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialiser le routeur Gin
	router := gin.Default()

	// Configurer les routes
	api.SetupRoutes(router)

	// Démarrer le serveur
	router.Run(":8080")
}
