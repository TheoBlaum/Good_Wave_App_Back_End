package handlers

import (
	"context"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"good_wave_back_end/models"
	"good_wave_back_end/database"
	"go.mongodb.org/mongo-driver/bson"
)

func ListSurfSpots(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.MongoDB.Collection("surfSpots").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des spots"})
		return
	}
	defer cursor.Close(ctx)

	var surfSpots []models.SurfSpot
	if err := cursor.All(ctx, &surfSpots); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du décodage des spots"})
		return
	}

	c.JSON(http.StatusOK, surfSpots)
}

func AddSurfSpot(c *gin.Context) {
	var spot models.SurfSpot

	if err := c.ShouldBindJSON(&spot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalide"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := database.MongoDB.Collection("surfSpots").InsertOne(ctx, spot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'insertion du spot"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Spot ajouté avec succès",
		"id":      result.InsertedID,
	})
}
