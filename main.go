package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/scottjbarr/yahoofinance"
)

// Display a usage message.
func usage() {
	fmt.Printf("Usage : %s symbol -config filename,yml\n", os.Args[0])
}

// Publish a new Quote to the Redis server.
func publish(quote *yahoofinance.Quote) {
	// log.Printf("%v", *quote)

	// create the Redis client
	c, err := Connect(&config.Redis)

	if c != nil {
		defer c.Close()
	}

	if err != nil {
		log.Printf("ERROR getting redis connection : %v", err)
		return
	}

	// format the JSON with a root element
	js, _ := json.Marshal(quote)
	price := fmt.Sprintf("{\"price\":%s}", js)

	c.Send("PUBLISH", config.Redis.Room, price)
	c.Flush()
}

// Repeatedly poll for Quote changes.
func poll(config *Config) {
	cache := make(map[string]yahoofinance.Quote)
	client := yahoofinance.CreateClient()

	for {
		// get quotes
		quotes, err := client.GetQuotes(config.Symbols)

		if err != nil {
			log.Printf("ERROR getting quote : %v", err)
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

	log.Printf("Loaded config file %v with config %+v", *configFile, *config)

	poll(config)
}
