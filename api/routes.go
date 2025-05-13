package api

import (
	"good_wave_back_end/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Route pour obtenir tous les spots de surf
	router.GET("/api/surf-spots", handlers.GetSurfSpots)
	// Route pour lister les spots de surf
	router.GET("/surf-spots", handlers.ListSurfSpots)
}
