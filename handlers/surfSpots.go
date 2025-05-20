package handlers

import (
	"context"
	"fmt"
	"good_wave_back_end/database"
	"good_wave_back_end/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cache     = make(map[string]interface{})
	cacheLock sync.RWMutex
	cacheTTL  = 5 * time.Minute
)

type PaginatedResponse struct {
	Data       []models.SurfSpot `json:"data"`
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
	TotalPages int               `json:"totalPages"`
	TotalItems int               `json:"totalItems"`
}

func GetSurfSpots(c *gin.Context) {
	// Récupérer les paramètres de pagination
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	forceRefresh := c.DefaultQuery("forceRefresh", "false")

	// Convertir les paramètres en entiers
	pageNum := 1
	pageSizeNum := 10
	fmt.Sscanf(page, "%d", &pageNum)
	fmt.Sscanf(pageSize, "%d", &pageSizeNum)

	// Calculer le skip pour MongoDB
	skip := (pageNum - 1) * pageSizeNum

	// Vérifier si les données sont en cache et si on ne force pas le rafraîchissement
	cacheKey := fmt.Sprintf("spots_%d_%d", pageNum, pageSizeNum)
	if forceRefresh != "true" {
		if data, ok := getFromCache(cacheKey); ok {
			c.JSON(http.StatusOK, data)
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Options de pagination pour MongoDB
	findOptions := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSizeNum))

	// Compter le nombre total de documents
	totalItems, err := database.MongoDB.Collection("surfSpots").CountDocuments(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du comptage des spots"})
		return
	}

	// Récupérer les spots avec pagination
	cursor, err := database.MongoDB.Collection("surfSpots").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des spots"})
		return
	}
	defer cursor.Close(ctx)

	var spots []models.SurfSpot
	if err := cursor.All(ctx, &spots); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du décodage des spots"})
		return
	}

	// Calculer le nombre total de pages
	totalPages := (int(totalItems) + pageSizeNum - 1) / pageSizeNum

	// Créer la réponse paginée
	response := PaginatedResponse{
		Data:       spots,
		Page:       pageNum,
		PageSize:   pageSizeNum,
		TotalPages: totalPages,
		TotalItems: int(totalItems),
	}

	// Mettre en cache la réponse
	setInCache(cacheKey, response)

	c.JSON(http.StatusOK, response)
}

func getFromCache(key string) (interface{}, bool) {
	cacheLock.RLock()
	defer cacheLock.RUnlock()

	if data, ok := cache[key]; ok {
		return data, true
	}
	return nil, false
}

func setInCache(key string, data interface{}) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	cache[key] = data

	// Nettoyer le cache après le TTL
	go func() {
		time.Sleep(cacheTTL)
		cacheLock.Lock()
		delete(cache, key)
		cacheLock.Unlock()
	}()
}

func RefreshCache(c *gin.Context) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	// Vider le cache
	cache = make(map[string]interface{})

	c.JSON(http.StatusOK, gin.H{"message": "Cache rafraîchi avec succès"})
}
