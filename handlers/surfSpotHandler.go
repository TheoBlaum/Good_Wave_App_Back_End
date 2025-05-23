package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"good_wave_back_end/database"
	"good_wave_back_end/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// S'assurer que l'ID est vide pour que MongoDB en génère un nouveau
	spot.ID = primitive.NilObjectID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := database.MongoDB.Collection("surfSpots").InsertOne(ctx, spot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'insertion du spot"})
		return
	}

	// Convertir l'ID inséré en ObjectID
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération de l'ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Spot ajouté avec succès",
		"id":      insertedID.Hex(),
	})
}

func UpdateSavedStatus(c *gin.Context) {
	var request struct {
		Saved bool `json:"saved"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Printf("Erreur de binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalide"})
		return
	}

	spotID := c.Param("id")
	if spotID == "" {
		fmt.Println("ID du spot manquant")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID du spot manquant"})
		return
	}

	fmt.Printf("Mise à jour du spot %s avec saved=%v\n", spotID, request.Saved)

	// Convertir l'ID string en ObjectID
	objectID, err := primitive.ObjectIDFromHex(spotID)
	if err != nil {
		fmt.Printf("Erreur de conversion de l'ID %s: %v\n", spotID, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Utiliser l'ObjectID dans le filtre
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"saved": request.Saved}}

	result, err := database.MongoDB.Collection("surfSpots").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Printf("Erreur de mise à jour MongoDB: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du spot"})
		return
	}

	if result.MatchedCount == 0 {
		fmt.Printf("Aucun spot trouvé avec l'ID %s\n", spotID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Spot non trouvé"})
		return
	}

	fmt.Printf("Spot mis à jour avec succès: %s\n", spotID)
	c.JSON(http.StatusOK, gin.H{"message": "Spot mis à jour avec succès"})
}
