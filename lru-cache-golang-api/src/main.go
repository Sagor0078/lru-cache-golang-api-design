package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Sagor0078/lru-cache-golang-api/src/cache"
    "github.com/Sagor0078/lru-cache-golang-api/src/handlers"
    "github.com/Sagor0078/lru-cache-golang-api/src/routes"
)

func main() {
    // Initialize the LRU cache
    lruCache := cache.NewLRUCache(100)
    handlers.Init(lruCache)

    // Create a Gin router
    router := gin.Default()

    // Setup routes
    routes.SetupRoutes(router)

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", router))
}