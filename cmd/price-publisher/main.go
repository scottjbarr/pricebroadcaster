package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/scottjbarr/pricebroadcaster"
	"github.com/scottjbarr/pricebroadcaster/pkg/cache"
	"github.com/scottjbarr/pricebroadcaster/pkg/clients"
	"github.com/scottjbarr/pricebroadcaster/pkg/models"
	"github.com/scottjbarr/redis"
)

// usage prints usage details
func usage() {
	fmt.Printf("Usage : %s symbol\n", os.Args[0])
}

func main() {
	cfg, err := pricebroadcaster.NewConfig()
	if err != nil {
		panic(err)
	}

	log.Printf("Starting with config %+v", cfg)

	bind := os.Getenv("BIND")
	if len(bind) == 0 {
		panic("BIND not specified")
	}

	pool, err := redis.NewPool(cfg.Redis.URL())
	if err != nil {
		panic(err)
	}

	cache := cache.NewRedisCache(pool)
	pub := pricebroadcaster.NewRedisPublisher(pool)

	broadcaster, err := pricebroadcaster.New(cfg.Room, pub, cache)
	if err != nil {
		panic(err)
	}

	ch := make(chan *models.OHLC, 1)

	// done := make(chan error, 1)

	go func() {
		broadcaster.Start(ch)
		// done <- err
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

// type Config struct {
// 	Bind     string `envconfig:"BIND"`
// 	RedisURL string `envconfig:"REDIS_URL"`
// }
