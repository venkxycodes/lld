package main

import (
	"fmt"
	"sync"
)

// Cache is a concurrency-safe, in-memory key-value cache.
// Create a Cache with NewCache before using it.
type Cache struct {
	mu    sync.RWMutex
	items map[string]string
}

func NewCache() *Cache {
	return &Cache{items: make(map[string]string)}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.items[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func main() {
	cache := NewCache()
	cache.Set("hello", "world")
	if value, ok := cache.Get("hello"); ok {
		fmt.Println(value)
	}
	cache.Delete("hello")
}
