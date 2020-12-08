package cache

import "sync"

// Use strings as cache keys
type Key string

// Can store any type of value in the cache
type Value interface{}

type Cache struct {
	data map[Key]Value
	lock sync.RWMutex
}

const Threaded = true

func (c *Cache) Get(k Key) (Value, bool) {
	if Threaded {
		// Use Read Lock/Unlock methods
		// since were aren't mutating the cache, only
		// reading values from it
		c.lock.RLock()
		defer c.lock.RUnlock()
	}

	value, exists := c.data[k]
	if !exists {
		return nil, false
	}
	return value, true
}

func (c *Cache) Set(k Key, v Value) {
	if Threaded {
		// Use the regular Lock/Unlock methods here
		// since we are mutating the cache
		c.lock.Lock()
		defer c.lock.Unlock()
	}
	c.data[k] = v
}

func New() *Cache {
	cache := &Cache {
		data: make(map[Key]Value),
	}
	return cache
}

func (c *Cache) Remove(k Key) {
	if Threaded {
		c.lock.Lock()
		defer c.lock.Unlock()
	}
	delete(c.data, k)
}