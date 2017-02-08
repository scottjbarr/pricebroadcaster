package pricebroadcaster

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/scottjbarr/yahoofinance"
)

// Broadcaster is responsible for receiving and broadcasting
// data
type Broadcaster struct {
	Config Config
	redis  redis.Conn
}

// New returns a new Broadcaster with the given Config
func New(config Config) (*Broadcaster, error) {
	conn, err := connect(config)

	if err != nil {
		return nil, fmt.Errorf("ERROR getting redis connection : %v", err)
	}

	return &Broadcaster{
		Config: config,
		redis:  conn,
	}, nil
}

// Start runs a Broadcaster
func (b Broadcaster) Start() {
	log.Printf("[Broadcaster] Starting with config %+v", b.Config)
	defer b.redis.Close()

	b.poll()
}

// polll for Quote changes, forever
func (b Broadcaster) poll() {
	cache := make(map[string]yahoofinance.Quote)
	client := yahoofinance.CreateClient()

	// duration to sleep for at the end of each loop
	sleep := time.Duration(b.Config.SleepTime) * time.Millisecond

	for {
		// get quotes
		quotes, err := client.GetQuotes(b.Config.Symbols)

		if err != nil {
			log.Printf("ERROR getting quotes : %v", err)
			continue
		}

		for _, quote := range quotes {
			if cache[quote.Symbol] != quote {
				// quote has changed, publish it
				cache[quote.Symbol] = quote
				b.publish(&quote)
			}
		}

		time.Sleep(sleep)
	}
}

// publish a new Quote to the Redis server.
func (b Broadcaster) publish(quote *yahoofinance.Quote) {
	// format the JSON with a root element
	js, _ := json.Marshal(quote)
	price := fmt.Sprintf("{\"price\":%s}", js)

	b.redis.Send("PUBLISH", b.Config.Redis.Room, price)
	b.redis.Flush()
}
