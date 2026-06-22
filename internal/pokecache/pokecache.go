package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{entries: make(map[string]cacheEntry)}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}
func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		cutoff := time.Now().Add(-interval)
		for key, entry := range c.entries {
			if entry.createdAt.Before(cutoff) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
