package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/scottjbarr/pricebroadcaster"
	"github.com/scottjbarr/pricebroadcaster/pkg/cache"
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

	bind := os.Getenv("BIND")
	if len(bind) == 0 {
		panic("BIND not specified")
	}

	pool, err := redis.NewPool(cfg.Redis.URL())
	if err != nil {
		panic(err)
	}

	cache := cache.NewRedisCache(pool)

	server := server{
		cache: cache,
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/prices/{symbol}", server.pricesHandler)
	if err := http.ListenAndServe(bind, r); err != nil {
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	//fmt.Fprintf(w, body)
	w.Write([]byte(body))
}
