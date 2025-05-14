package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

type CreateSurfSpotRequest struct {
	Fields struct {
		SurfBreak        []string `json:"Surf Break"`
		DifficultyLevel  int      `json:"Difficulty Level"`
		Destination      string   `json:"Destination"`
		Geocode          string   `json:"Geocode"`
		MagicSeaweedLink string   `json:"Magic Seaweed Link"`
		Photos           []struct {
			URL string `json:"url"`
		} `json:"Photos"`
		PeakSurfSeasonBegins    string `json:"Peak Surf Season Begins"`
		DestinationStateCountry string `json:"Destination State/Country"`
		PeakSurfSeasonEnds      string `json:"Peak Surf Season Ends"`
	} `json:"fields"`
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

func convertToSurfSpot(request CreateSurfSpotRequest) SurfSpot {
	// Créer les thumbnails par défaut
	defaultThumbnail := struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	}{
		URL:    request.Fields.Photos[0].URL,
		Width:  800,
		Height: 600,
	}

	// Convertir les photos
	photos := make([]struct {
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
	}, len(request.Fields.Photos))

	for i, photo := range request.Fields.Photos {
		photos[i] = struct {
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
		}{
			ID:       time.Now().Format("20060102150405") + "_" + string(rune(i)),
			URL:      photo.URL,
			Filename: "photo_" + time.Now().Format("20060102150405") + "_" + string(rune(i)) + ".jpg",
			Size:     0,
			Type:     "image/jpeg",
			Thumbnails: struct {
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
			}{
				Small: defaultThumbnail,
				Large: defaultThumbnail,
				Full:  defaultThumbnail,
			},
		}
	}

	return SurfSpot{
		ID: time.Now().Format("20060102150405"),
		Fields: struct {
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
		}{
			SurfBreak:               request.Fields.SurfBreak,
			DifficultyLevel:         request.Fields.DifficultyLevel,
			Destination:             request.Fields.Destination,
			Geocode:                 request.Fields.Geocode,
			Influencers:             []string{},
			MagicSeaweedLink:        request.Fields.MagicSeaweedLink,
			Photos:                  photos,
			PeakSurfSeasonBegins:    request.Fields.PeakSurfSeasonBegins,
			DestinationStateCountry: request.Fields.DestinationStateCountry,
			PeakSurfSeasonEnds:      request.Fields.PeakSurfSeasonEnds,
			Address:                 "",
		},
		CreatedTime: time.Now().Format(time.RFC3339),
	}
}

func AddSurfSpot(c *gin.Context) {
	var request CreateSurfSpotRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Données invalides",
		})
		return
	}

	// Lire le fichier JSON existant
	jsonFile := filepath.Join("data", "surfSpots.json")
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de la lecture du fichier JSON",
		})
		return
	}

	// Décoder le JSON existant
	var response SurfSpotsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors du décodage du JSON",
		})
		return
	}

	// Convertir la requête en SurfSpot
	newSpot := convertToSurfSpot(request)

	// Ajouter le nouveau spot
	response.Records = append(response.Records, newSpot)

	// Encoder et sauvegarder le nouveau JSON
	newData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de l'encodage du JSON",
		})
		return
	}

	if err := os.WriteFile(jsonFile, newData, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de l'écriture du fichier JSON",
		})
		return
	}

	c.JSON(http.StatusCreated, newSpot)
}
