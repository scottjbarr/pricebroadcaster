package cache

import "errors"

var ErrCacheMiss = errors.New("Cache miss")

type Cache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type NoopCache struct {
}

func NewNoopCache() *NoopCache {
	return &NoopCache{}
}

func (c *NoopCache) Set(key string, payload []byte) error {
	return nil
}

func (c *NoopCache) Get(key string) ([]byte, error) {
	return nil, nil
}
