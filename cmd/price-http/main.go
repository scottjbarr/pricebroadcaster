package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/scottjbarr/config"
	"github.com/scottjbarr/pricebroadcaster/pkg/cache"
	"github.com/scottjbarr/redis"
)

// usage prints usage details
func usage() {
	fmt.Printf("Usage : %s symbol\n", os.Args[0])
}

func main() {
	cfg := Config{}
	if err := config.Process(&cfg); err != nil {
		panic(err)
	}

	pool, err := redis.NewPool(cfg.RedisURL)
	if err != nil {
		panic(err)
	}

	cache := cache.NewRedisCache(pool)

	server := server{
		cache: cache,
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/prices/{symbol}", server.pricesHandler)
	if err := http.ListenAndServe(cfg.Bind, r); err != nil {
		panic(err)
	}
}

type server struct {
	cache cache.Cache
}

func (s server) pricesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	symbol := vars["symbol"]

	body, err := s.cache.Get(symbol)
	if err != nil {
		status := http.StatusInternalServerError
		if err == cache.ErrCacheMiss {
			status = http.StatusNotFound
		}

		w.WriteHeader(status)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(body))
}

type Config struct {
	RedisURL string `envconfig:"REDIS_URL" required:"true"`
	Room     string `envconfig:"ROOM" required:"true"`
	Bind     string `envconfig:"BIND" required:"true"`
}
