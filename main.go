package main

import (
	//"context"
	"good_wave_back_end/api"
	"good_wave_back_end/database"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Charger les variables d'environnement
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	// Récupérer les variables d'environnement
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		log.Fatal("La variable d'environnement MONGODB_URI n'est pas définie")
	}

	dbName := os.Getenv("MONGODB_DB_NAME")
	if dbName == "" {
		log.Fatal("La variable d'environnement MONGODB_DB_NAME n'est pas définie")
	}

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

	err := database.ConnectWithOptions(mongodbURI, dbName, serverAPI)
	if err != nil {
		log.Fatal("Erreur de connexion à MongoDB: ", err)
	}

	// Configurer les routes
	api.SetupRoutes(router)

	// Démarrer le serveur
	router.Run(":8080")
}
