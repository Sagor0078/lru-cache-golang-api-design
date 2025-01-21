package handlers

import (
    "net/http"
    "sync"

    "github.com/gin-gonic/gin"
    "github.com/Sagor0078/lru-cache-golang-api/src/cache"
)

var (
    mu    sync.Mutex
    lru   *cache.LRUCache
)

func Init(c *cache.LRUCache) {
    lru = c
}

func AddCacheEntry(c *gin.Context) {
    var entry cache.CacheEntry
    if err := c.ShouldBindJSON(&entry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    mu.Lock()
    defer mu.Unlock()

    lru.Add(entry.Key, entry.Value)
    c.Status(http.StatusCreated)
}

func GetCacheEntry(c *gin.Context) {
    key := c.Param("key")

    mu.Lock()
    defer mu.Unlock()

    value, found := lru.Get(key)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "key":   key,
        "value": value,
    })
}