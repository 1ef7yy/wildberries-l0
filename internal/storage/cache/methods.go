package cache

import "encoding/json"

func (c *Cache) Set(key string, value json.RawMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get retrieves the value associated with the given key from the cache.
func (c *Cache) Get(key string) (json.RawMessage, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]json.RawMessage)
}
