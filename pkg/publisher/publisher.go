package publisher

import (
	redigo "github.com/gomodule/redigo/redis"
)

const (
	cmdRedisPublish = "PUBLISH"
)

type Publisher interface {
	Publish([]byte) error
}

type RedisPublisher struct {
	Pool *redigo.Pool
	Key  string
}

func NewRedisPublisher(pool *redigo.Pool, key string) *RedisPublisher {
	return &RedisPublisher{
		Pool: pool,
		Key:  key,
	}
}

func (p *RedisPublisher) Publish(b []byte) error {
	conn := p.Pool.Get()
	defer conn.Close()

	if _, err := conn.Do(cmdRedisPublish, p.Key, b); err != nil {
		return err
	}

	if err := conn.Flush(); err != nil {
		return err
	}

	return nil
}
