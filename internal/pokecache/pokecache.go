package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	Entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.Entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.Entries[key]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		staleEntries := []string{}
		for k, v := range c.Entries {
			if time.Since(v.createdAt) > c.interval {
				staleEntries = append(staleEntries, k)
			}
		}
		for _, key := range staleEntries {
			delete(c.Entries, key)
		}
		c.mu.Unlock()
	}

}
