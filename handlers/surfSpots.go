package handlers

import (
	"context"
	"time"
	"net/http"
	"good_wave_back_end/database"
	"good_wave_back_end/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSurfSpots(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.MongoDB.Collection("surfSpots").Find(ctx, bson.M{})
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

	fmt.Println("📦 Spots récupérés :", spots)
	c.JSON(http.StatusOK, spots)
}
