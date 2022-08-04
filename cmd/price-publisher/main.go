package main

import (
	"log"
	"time"

	"github.com/scottjbarr/config"
	"github.com/scottjbarr/pricebroadcaster/pkg/broadcaster"
	"github.com/scottjbarr/pricebroadcaster/pkg/cache"
	"github.com/scottjbarr/pricebroadcaster/pkg/clients"
	"github.com/scottjbarr/pricebroadcaster/pkg/models"
	"github.com/scottjbarr/pricebroadcaster/pkg/publisher"
	"github.com/scottjbarr/redis"
)

func main() {
	cfg := Config{}
	if err := config.Process(&cfg); err != nil {
		panic(err)
	}

	log.Printf("Starting with config %+v", cfg)

	pool, err := redis.NewPool(cfg.RedisURL)
	if err != nil {
		panic(err)
	}

	cache := cache.NewRedisCache(pool)
	pub := publisher.NewRedisPublisher(pool)

	broadcaster, err := broadcaster.New(cfg.Room, pub, cache)
	if err != nil {
		panic(err)
	}

	ch := make(chan *models.OHLC, 1)

	go func() {
		broadcaster.Start(ch)
	}()

	// start fetching and pushing prices
	c := clients.NewMockClient()
	t := time.NewTicker(time.Second * 10)

	for {
		<-t.C
		ohlc, err := c.Get("whatever")
		if err != nil {
			panic(err)
		}

		ch <- ohlc
	}
}

type Config struct {
	RedisURL string `envconfig:"REDIS_URL"`
	Room     string `envconfig:"ROOM"`
}
