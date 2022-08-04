package cache

import (
	"sync"
)

type MemoryCache struct {
	store map[string][]byte
	mutex *sync.Mutex
}

func NewMemoryCache() MemoryCache {
	return MemoryCache{
		store: make(map[string][]byte),
		mutex: &sync.Mutex{},
	}
}

func (m MemoryCache) Set(key string, payload []byte) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.store[key] = payload

	return nil
}

func (m MemoryCache) Get(key string) ([]byte, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data, ok := m.store[key]
	if !ok {
		return nil, ErrCacheMiss
	}

	return data, nil
}
