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

func NewCache(interval time.Duration) (*Cache, error) {
	c := Cache{}
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
				c.ReapLoop(interval)
		}
	}()

	return &c, nil
}

func (c *Cache) Add(url string, val []byte, mutex *sync.RWMutex) error {
	fmt.Println("saving locations to cache...")
	ce := CacheEntry{}
	ce.Value = val
	ce.CreatedAt = time.Now()
	mutex.Lock()
	defer mutex.Unlock()
	(*c)[url] = ce
	
	return nil
}

func (c *Cache) Get(url string, mutex *sync.RWMutex) ([]byte, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return (*c)[url].Value, nil
}

func(c *Cache) ReapLoop(interval time.Duration) {
	for k, v := range (*c) {
		if v.CreatedAt.Before(time.Now().Add(-interval)) {
			delete(*c, k)
		}
	}
}
