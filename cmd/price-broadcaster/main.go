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
	"github.com/scottjbarr/pricebroadcaster/pkg/service"
	"github.com/scottjbarr/redis"
)

func main() {
	cfg := Config{}
	if err := config.Process(&cfg); err != nil {
		panic(err)
	}

	log.Printf("Starting with config %+v", cfg)

	pool, err := redis.NewPoolWithBorrowFunc(cfg.RedisURL, redis.NoopOnBorrow)
	if err != nil {
		panic(err)
	}

	// build the minor dependencies
	cache := cache.NewRedisCache(pool)
	pub := publisher.NewRedisPublisher(pool, cfg.Room)
	broadcaster := broadcaster.NewWithCache(pub, cache)

	// the service that brings it all together
	service := service.New(broadcaster)

	ch := make(chan *models.OHLC, 1)

	// start fetching and pushing prices
	c := clients.NewMockClient()
	t := time.NewTicker(time.Second * 10)

	go func() {
		for {
			<-t.C
			ohlc, err := c.Get("whatever")
			if err != nil {
				panic(err)
			}

			ch <- ohlc
		}
	}()

	if err := service.Run(ch); err != nil {
		panic(err)
	}
}

type Config struct {
	RedisURL string `envconfig:"REDIS_URL"`
	Room     string `envconfig:"ROOM"`
}
