package main

import (
	//"context"
	"log"
	"good_wave_back_end/api"
	"good_wave_back_end/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// Connexion à MongoDB selon la documentation
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := "mongodb+srv://admin:admin@goodwave.db0uaj7.mongodb.net/?retryWrites=true&w=majority&appName=GoodWave"
	
	err := database.ConnectWithOptions(uri, "GoodWave", serverAPI)
	if err != nil {
		log.Fatal("Erreur de connexion à MongoDB: ", err)
	}

	// Configurer les routes
	api.SetupRoutes(router)

	// Démarrer le serveur
	router.Run(":8080")
}
