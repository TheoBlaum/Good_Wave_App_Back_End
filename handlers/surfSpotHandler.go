package handlers

import (
	"context"
	"net/http"
	"time"

	"good_wave_back_end/database"
	"good_wave_back_end/models"

	"github.com/gin-gonic/gin"
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

func UpdateSavedStatus(c *gin.Context) {
	var request struct {
		Saved bool `json:"saved"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalide"})
		return
	}

	spotID := c.Param("id")
	if spotID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID du spot manquant"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Utiliser directement l'ID comme string au lieu de le convertir en ObjectID
	filter := bson.M{"_id": spotID}
	update := bson.M{"$set": bson.M{"saved": request.Saved}}

	_, err := database.MongoDB.Collection("surfSpots").UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du spot"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Spot mis à jour avec succès"})
}
