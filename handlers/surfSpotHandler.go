package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"good_wave_back_end/models"

	"github.com/gin-gonic/gin"
)

type SurfSpotResponse struct {
	Records []struct {
		Fields struct {
			Photos []struct {
				URL string `json:"url"`
			} `json:"Photos"`
			Title            string   `json:"Destination"`
			Location         string   `json:"Destination State/Country"`
			PeakSeasonBegins string   `json:"Peak Surf Season Begins"`
			PeakSeasonEnds   string   `json:"Peak Surf Season Ends"`
			SurfBreak        []string `json:"Surf Break"`
			DifficultyLevel  int      `json:"Difficulty Level"`
			ForecastURL		string		`json:"Magic Seaweed Link"`
		} `json:"fields"`
	} `json:"records"`
}

func ListSurfSpots(c *gin.Context) {
	// Lire le fichier JSON
	jsonFile := filepath.Join("data", "surfSpots.json")
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture du fichier"})
		return
	}

	// Décoder le JSON
	var response SurfSpotResponse
	if err := json.Unmarshal(data, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du décodage du JSON"})
		return
	}

	// Transformer les données
	var spots []models.SurfSpotList
	for _, record := range response.Records {
		spot := models.SurfSpotList{
			Title:            record.Fields.Title,
			Location:         record.Fields.Location,
			PeakSeasonBegins: record.Fields.PeakSeasonBegins,
			PeakSeasonEnds:   record.Fields.PeakSeasonEnds,
			SurfBreak:        record.Fields.SurfBreak,
			DifficultyLevel:  record.Fields.DifficultyLevel,
			ForecastURL:  	  record.Fields.ForecastURL,
		}

		// Ajouter l'URL de la photo si elle existe
		if len(record.Fields.Photos) > 0 {
			spot.PhotoURL = record.Fields.Photos[0].URL
		}

		spots = append(spots, spot)
	}

	c.JSON(http.StatusOK, spots)
}
