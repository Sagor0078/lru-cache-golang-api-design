package cache

import (
    "container/list"
    "sync"
)

type CacheEntry struct {
    Key   string
    Value interface{}
}

type LRUCache struct {
    capacity int
    cache    map[string]*list.Element
    lruList  *list.List
    mu       sync.Mutex
}

type entry struct {
    key   string
    value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[string]*list.Element),
        lruList:  list.New(),
    }
}

func (c *LRUCache) Add(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()

    if el, ok := c.cache[key]; ok {
        c.lruList.MoveToFront(el)
        el.Value.(*entry).value = value
        return
    }

    if c.lruList.Len() == c.capacity {
        el := c.lruList.Back()
        if el != nil {
            c.lruList.Remove(el)
            delete(c.cache, el.Value.(*entry).key)
        }
    }

    el := c.lruList.PushFront(&entry{key, value})
    c.cache[key] = el
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    if el, ok := c.cache[key]; ok {
        c.lruList.MoveToFront(el)
        return el.Value.(*entry).value, true
    }
    return nil, false
}