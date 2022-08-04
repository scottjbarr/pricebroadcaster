package cache

import (
	redigo "github.com/gomodule/redigo/redis"
)

type RedisCache struct {
	Pool *redigo.Pool
}

func NewRedisCache(pool *redigo.Pool) *RedisCache {
	return &RedisCache{
		Pool: pool,
	}
}

func (c *RedisCache) Set(key string, value []byte) error {
	conn := c.Pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", key, value); err != nil {
		return err
	}

	return nil
}

func (c *RedisCache) Get(key string) ([]byte, error) {
	conn := c.Pool.Get()
	defer conn.Close()

	resp, err := redigo.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
