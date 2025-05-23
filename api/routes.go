package api

import (
	"good_wave_back_end/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/surf-spots", handlers.GetSurfSpots)
	router.GET("/surf-spots", handlers.ListSurfSpots)
	router.POST("/api/surf-spots", handlers.AddSurfSpot)
	router.PUT("/api/surf-spots/:id", handlers.UpdateSavedStatus)
	router.POST("/api/refresh-cache", handlers.RefreshCache)
}
