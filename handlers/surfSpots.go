package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type SurfSpot struct {
	ID     string `json:"id"`
	Fields struct {
		SurfBreak        []string `json:"Surf Break"`
		DifficultyLevel  int      `json:"Difficulty Level"`
		Destination      string   `json:"Destination"`
		Geocode          string   `json:"Geocode"`
		Influencers      []string `json:"Influencers"`
		MagicSeaweedLink string   `json:"Magic Seaweed Link"`
		Photos           []struct {
			ID         string `json:"id"`
			URL        string `json:"url"`
			Filename   string `json:"filename"`
			Size       int    `json:"size"`
			Type       string `json:"type"`
			Thumbnails struct {
				Small struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"small"`
				Large struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"large"`
				Full struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"full"`
			} `json:"thumbnails"`
		} `json:"Photos"`
		PeakSurfSeasonBegins    string `json:"Peak Surf Season Begins"`
		DestinationStateCountry string `json:"Destination State/Country"`
		PeakSurfSeasonEnds      string `json:"Peak Surf Season Ends"`
		Address                 string `json:"Address"`
	} `json:"fields"`
	CreatedTime string `json:"createdTime"`
}

type SurfSpotsResponse struct {
	Records []SurfSpot `json:"records"`
	Offset  string     `json:"offset"`
}

func GetSurfSpots(c *gin.Context) {
	// Lire le fichier JSON
	jsonFile, err := os.ReadFile(filepath.Join("data", "surfSpots.json"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de la lecture du fichier JSON",
		})
		return
	}

	// Définir le type de contenu de la réponse
	c.Header("Content-Type", "application/json")

	// Écrire le contenu JSON directement dans la réponse
	c.Data(http.StatusOK, "application/json", jsonFile)
}
