package pricebroadcaster

import (
	"encoding/json"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/scottjbarr/pricebroadcaster/pkg/cache"
	"github.com/scottjbarr/pricebroadcaster/pkg/models"
)

// Broadcaster is responsible for receiving and broadcasting
// data
type Broadcaster struct {
	Room      string
	Publisher Publisher
	Cache     cache.Cache
}

// New returns a new Broadcaster with the given Config
func New(room string, pub Publisher, cache cache.Cache) (*Broadcaster, error) {
	return &Broadcaster{
		Room:      room,
		Publisher: pub,
		Cache:     cache,
	}, nil
}

// Start runs a Broadcaster
func (b Broadcaster) Start(ch chan *models.OHLC) {
	for {
		select {
		case ohlc := <-ch:
			if err := b.publish(ohlc); err != nil {
				fmt.Printf("ERROR %v\n", err)
			}
		}
	}
}

// publish a new Quote to the Redis server.
func (b Broadcaster) publish(ohlc *models.OHLC) error {
	// format the JSON with a root element
	data, err := json.Marshal(ohlc)
	if err != nil {
		return err
	}

	if err := b.Publisher.Publish(b.Room, data); err != nil {
		return err
	}

	return b.Cache.Set(ohlc.Symbol, data)
}

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
