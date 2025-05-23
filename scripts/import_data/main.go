package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"good_wave_back_end/database"
	"good_wave_back_end/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ImportRecord struct {
	ID          string                 `json:"id"`
	Fields      map[string]interface{} `json:"fields"`
	CreatedTime string                 `json:"createdTime"`
}

type ImportData struct {
	Records []ImportRecord `json:"records"`
	Offset  string         `json:"offset"`
}

func main() {
	// Charger les variables d'environnement
	if err := godotenv.Load(filepath.Join("..", "..", ".env")); err != nil {
		fmt.Printf("Erreur lors du chargement du fichier .env: %v\n", err)
		os.Exit(1)
	}

	// Récupérer les variables d'environnement
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		fmt.Println("La variable d'environnement MONGODB_URI n'est pas définie")
		os.Exit(1)
	}

	dbName := os.Getenv("MONGODB_DB_NAME")
	if dbName == "" {
		fmt.Println("La variable d'environnement MONGODB_DB_NAME n'est pas définie")
		os.Exit(1)
	}

	// Connexion à MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	if err := database.ConnectWithOptions(mongodbURI, dbName, serverAPI); err != nil {
		fmt.Printf("Erreur de connexion à MongoDB: %v\n", err)
		os.Exit(1)
	}

	// Lire le fichier JSON converti
	jsonFile := filepath.Join("..", "..", "data", "surfSpots_converted.json")
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier JSON: %v\n", err)
		os.Exit(1)
	}

	// Décoder le JSON
	var data ImportData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("Erreur lors du décodage du JSON: %v\n", err)
		os.Exit(1)
	}

	// Créer un contexte avec timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Vider la collection existante
	if err := database.MongoDB.Collection("surfSpots").Drop(ctx); err != nil {
		fmt.Printf("Erreur lors de la suppression de la collection: %v\n", err)
		os.Exit(1)
	}

	// Convertir et insérer les données
	for _, record := range data.Records {
		// Convertir l'ID string en ObjectID
		objectID, err := primitive.ObjectIDFromHex(record.ID)
		if err != nil {
			fmt.Printf("Erreur lors de la conversion de l'ID %s: %v\n", record.ID, err)
			continue
		}

		// Vérifier si le spot existe déjà dans la base de données
		var existingSpot models.SurfSpot
		err = database.MongoDB.Collection("surfSpots").FindOne(ctx, bson.M{"destination": record.Fields["Destination"].(string)}).Decode(&existingSpot)
		saved := false
		if err == nil {
			// Si le spot existe, préserver son état saved
			saved = existingSpot.Saved
		}

		// Créer un nouveau spot
		spot := models.SurfSpot{
			ID:          objectID,
			Destination: record.Fields["Destination"].(string),
			Address:     record.Fields["Address"].(string),
			Country:     record.Fields["Destination State/Country"].(string),
			Difficulty:  int(record.Fields["Difficulty Level"].(float64)),
			SurfBreak:   convertToStringSlice(record.Fields["Surf Break"].([]interface{})),
			SeasonStart: record.Fields["Peak Surf Season Begins"].(string),
			SeasonEnd:   record.Fields["Peak Surf Season Ends"].(string),
			Photo:       getPhotoURL(record.Fields["Photos"].([]interface{})),
			Link:        record.Fields["Magic Seaweed Link"].(string),
			Geocode:     record.Fields["Geocode"].(string),
			Saved:       saved,
		}

		// Insérer le spot dans MongoDB
		_, err = database.MongoDB.Collection("surfSpots").InsertOne(ctx, spot)
		if err != nil {
			fmt.Printf("Erreur lors de l'insertion du spot %s: %v\n", spot.Destination, err)
			continue
		}
	}

	fmt.Println("Import des données terminé avec succès")
}

func convertToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = v.(string)
	}
	return result
}

func getPhotoURL(photos []interface{}) string {
	if len(photos) == 0 {
		return ""
	}
	photo := photos[0].(map[string]interface{})
	return photo["url"].(string)
}
