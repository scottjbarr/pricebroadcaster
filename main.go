package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/scottjbarr/yahoofinance"
)

// Display a usage message.
func usage() {
	fmt.Printf("Usage : %s symbol -config filename,yml\n", os.Args[0])
}

// Publish a new Quote to the Redis server.
func publish(quote *yahoofinance.Quote) {
	js, _ := json.Marshal(quote)

	// format the JSON with a root element
	price := fmt.Sprintf("{price:%s}", js)

	redisClient.Send("PUBLISH", config.Redis.Room, price)
	redisClient.Flush()
}

// Repeatedly poll for Quote changes.
func poll(config *Config) {
	cache := make(map[string]yahoofinance.Quote)
	client := yahoofinance.CreateClient()

	for true {
		// get quotes
		quotes, err := client.GetQuotes(config.Symbols)

		if err != nil {
			log.Printf("ERROR : %v", err)
			continue
		}

		for _, quote := range quotes {
			if cache[quote.Symbol] != quote {
				// quote has changed, publish it
				cache[quote.Symbol] = quote
				publish(&quote)
			}
		}

		time.Sleep(time.Duration(config.SleepTime) * time.Millisecond)
	}
}

var redisClient redis.Conn
var config *Config

func main() {
	configFile := flag.String("config", "", "Config file")

	flag.Parse()

	if *configFile == "" {
		fmt.Println("A -config must be given. See config/example.conf.sample")
		os.Exit(1)
	}

	// parse the config file
	config = ParseConfig(*configFile)

	f, err := os.OpenFile(
		config.LogFile,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer f.Close()

	// send log output to the log file
	log.SetOutput(f)

	log.Printf("Loaded config file %v with config %v", *configFile, *config)

	// create the Redis client
	redisClient = Connect(&config.Redis)
	defer redisClient.Close()

	poll(config)
}
