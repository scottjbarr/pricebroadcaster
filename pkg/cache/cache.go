package cache

import "errors"

var ErrCacheMiss = errors.New("Cache miss")

type Cache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}
