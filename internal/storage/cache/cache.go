package cache

import (
	"encoding/json"
	"sync"
)

type Cache struct {
	data map[string]json.RawMessage
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]json.RawMessage),
		mu:   sync.RWMutex{},
	}
}
