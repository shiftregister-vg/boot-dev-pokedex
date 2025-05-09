package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	ttl     time.Duration
	mut     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mut.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var (
		val   []byte
		found bool = false
	)

	c.mut.Lock()

	if entry, ok := c.entries[key]; ok {
		found = true
		val = entry.val
	}

	c.mut.Unlock()

	return val, found
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)

	select {
	case <-ticker.C:
		c.mut.Lock()
		for key, entry := range c.entries {
			elapsed := time.Since(entry.createdAt)
			if elapsed >= c.ttl {
				delete(c.entries, key)
			}
		}
		c.mut.Unlock()
	}
}

func NewCache(ttl time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		ttl:     ttl,
		mut:     sync.Mutex{},
	}

	go cache.reapLoop()

	return &cache
}
