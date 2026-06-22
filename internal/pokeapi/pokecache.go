package pokeapi

import (
	"sync"
	"time"
	"fmt"
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
	entry, err := c.entries[key]
	if err != nil {
		fmt.Println("error fetching data")
	}
	return entry.val, ok
}
func (c *Cache) reapLoop(interval) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ticker = time.NewTicker(interval)
	cutoff := time.Now().Add(-interval)
	for key, entry := range c.entries {
		if entry.createdAt.Before(cutoff) {
			delete(c.entries, key)
		}
	}
}
