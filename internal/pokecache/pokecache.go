package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
	tick     time.Time
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		mutex:    sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) (val []byte, entryfound bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for stringKey, cacheEntryVal := range cache.entries {
		if stringKey == key {
			val = cacheEntryVal.value
			entryfound = true
		}
		break
	}
	return val, entryfound
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		cache.mutex.Lock()
		for key, entry := range cache.entries {
			if time.Since(entry.createdAt) > cache.interval {
				delete(cache.entries, key)
			}
		}
		cache.mutex.Unlock()
	}
}
