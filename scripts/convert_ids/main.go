package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID          string                 `json:"id"`
	Fields      map[string]interface{} `json:"fields"`
	CreatedTime string                 `json:"createdTime"`
}

type SurfSpotsData struct {
	Records []Record `json:"records"`
	Offset  string   `json:"offset"`
}

func main() {
	// Charger les variables d'environnement
	if err := godotenv.Load(filepath.Join("..", "..", ".env")); err != nil {
		fmt.Printf("Erreur lors du chargement du fichier .env: %v\n", err)
		os.Exit(1)
	}

	// Lire le fichier JSON
	jsonFile := filepath.Join("..", "..", "data", "surfSpots.json")
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier JSON: %v\n", err)
		os.Exit(1)
	}

	// Décoder le JSON
	var data SurfSpotsData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("Erreur lors du décodage du JSON: %v\n", err)
		os.Exit(1)
	}

	// Convertir les IDs
	for i := range data.Records {
		// Générer un nouvel ObjectID
		newID := primitive.NewObjectID()
		// Mettre à jour l'ID dans les données
		data.Records[i].ID = newID.Hex()
	}

	// Encoder les données modifiées
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Erreur lors de l'encodage du JSON: %v\n", err)
		os.Exit(1)
	}

	// Écrire le fichier mis à jour
	outputFile := filepath.Join("..", "..", "data", "surfSpots_converted.json")
	if err := ioutil.WriteFile(outputFile, updatedJSON, 0644); err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Conversion terminée avec succès. Le fichier converti a été sauvegardé dans data/surfSpots_converted.json")
}
