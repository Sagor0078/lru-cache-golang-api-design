package cache

import (
    "testing"
)

func TestCache(t *testing.T) {
    cache := NewCache(2)

    cache.Add(1, 100)
    value, ok := cache.Get(1)
    if !ok || value != 100 {
        t.Errorf("Expected 100, got %v", value)
    }

    cache.Add(2, 200)
    value, ok = cache.Get(2)
    if !ok || value != 200 {
        t.Errorf("Expected 200, got %v", value)
    }

    cache.Add(3, 300) // This should evict key 1
    _, ok = cache.Get(1)
    if ok {
        t.Error("Expected key 1 to be evicted")
    }

    value, ok = cache.Get(2)
    if !ok || value != 200 {
        t.Errorf("Expected 200, got %v", value)
    }

    value, ok = cache.Get(3)
    if !ok || value != 300 {
        t.Errorf("Expected 300, got %v", value)
    }
}