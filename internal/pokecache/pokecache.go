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

func NewCache(interval time.Duration) Cache {
	return Cache{
		interval: interval,
	}
}

// TODO:
func (c *Cache) Add(key string, val []byte) {
	// ...
}

// TODO:
func (c *Cache) Get(key string) ([]byte, bool) {
	// ...
	return []byte{}, false
}

// TODO:
func (c *Cache) reapLoop() {

}
