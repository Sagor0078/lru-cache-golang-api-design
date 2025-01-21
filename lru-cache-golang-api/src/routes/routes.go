package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Sagor0078/lru-cache-golang-api/src/handlers"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/cache", handlers.AddCacheEntry)
    router.GET("/cache/:key", handlers.GetCacheEntry)
}