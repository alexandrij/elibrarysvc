package cache

import (
	"errors"
	"sync"
	"time"
)

var ErrItemNotFound = errors.New("cache: item not found")

type item struct {
	value     interface{}
	createdAt int64
	ttl       int64
}

type MemoryCache struct {
	cache map[interface{}]*item
	mtx   sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{cache: make(map[interface{}]*item)}
	go c.setTtlTimer()
	return c
}

func (c *MemoryCache) setTtlTimer() {
	for {
		c.mtx.Lock()
		for k, v := range c.cache {
			if time.Now().Unix()-v.createdAt > v.ttl {
				delete(c.cache, k)
			}
		}
		c.mtx.Unlock()

		<-time.After(time.Second)
	}
}

func (c *MemoryCache) Set(key, value interface{}, ttl int64) error {
	c.mtx.Lock()
	c.cache[key] = &item{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       ttl,
	}
	c.mtx.Unlock()

	return nil
}

func (c *MemoryCache) Get(key interface{}) (interface{}, error) {
	c.mtx.RLock()
	item, ok := c.cache[key]
	c.mtx.RUnlock()

	if !ok {
		return nil, ErrItemNotFound
	}

	return item.value, nil
}
