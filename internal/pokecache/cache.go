package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mx       sync.RWMutex
	data     map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mx.Lock()
	defer c.mx.Unlock()

	copied := make([]byte, len(val))
	copy(copied, val)

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       copied,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	data, ok := c.data[key]
	if !ok {
		return nil, false
	}
	
	copied := make([]byte, len(data.val))
	copy(copied, data.val)

	return copied, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for t := range ticker.C {
		c.mx.Lock()
		for key, entry := range c.data {
			if t.Sub(entry.createdAt) > c.interval {
				delete(c.data, key)
			}
		}
		c.mx.Unlock()
	}
}
