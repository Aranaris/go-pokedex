package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Value []byte
}

type Cache map[string]CacheEntry

func NewCache(interval time.Duration) (Cache, error) {
	c := make(map[string]CacheEntry)
	return c, nil
}

func (c *Cache) Add(url string, val []byte, mutex *sync.RWMutex) error {
	fmt.Println("saving to cache")
	ce := CacheEntry{}
	ce.Value = val
	ce.CreatedAt = time.Now()
	mutex.Lock()
	(*c)[url] = ce
	mutex.Unlock()
	return nil
}
