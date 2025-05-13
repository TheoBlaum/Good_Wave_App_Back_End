package main

import (
	"good_wave_back_end/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialiser le routeur Gin
	router := gin.Default()

	// Configuration CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // En production, spécifiez les domaines autorisés
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 heures
	}))

	// Configurer les routes
	api.SetupRoutes(router)

	// Démarrer le serveur
	router.Run(":8080")
}
