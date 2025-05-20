package api

import (
	"good_wave_back_end/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Route pour obtenir tous les spots de surf avec pagination
	router.GET("/api/surf-spots", handlers.GetSurfSpots)
	// Route pour lister les spots de surf
	router.GET("/surf-spots", handlers.ListSurfSpots)
	// Route pour ajouter un nouveau spot de surf
	router.POST("/api/surf-spots", handlers.AddSurfSpot)
	// Route pour ajouter ou enlever un spot de la liste des favoris
	router.PUT("/api/surf-spots/:id", handlers.UpdateSavedStatus)
	// Route pour rafraîchir le cache
	router.POST("/api/refresh-cache", handlers.RefreshCache)
}
