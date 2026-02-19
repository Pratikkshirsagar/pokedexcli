package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := Cache{}
	c.cache = make(map[string]cacheEntry)

	go c.reapLoop(interval)
	return &c
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	dat, exist := c.cache[key]

	if !exist {
		return nil, false
	}

	return dat.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for name, entry := range c.cache {
			age := time.Since(entry.createdAt)

			if age > interval {
				delete(c.cache, name)
			}
		}
		c.mu.Unlock()
	}

}
