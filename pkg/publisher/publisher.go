package publisher

import (
	redigo "github.com/gomodule/redigo/redis"
)

type Publisher interface {
	Publish(string, []byte) error
}

type RedisPublisher struct {
	Pool *redigo.Pool
}

func NewRedisPublisher(pool *redigo.Pool) *RedisPublisher {
	return &RedisPublisher{
		Pool: pool,
	}
}

func (p *RedisPublisher) Publish(key string, b []byte) error {
	conn := p.Pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PUBLISH", key, b); err != nil {
		return err
	}

	if err := conn.Flush(); err != nil {
		return err
	}

	return nil
}
