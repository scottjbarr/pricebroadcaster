package broadcaster

import (
	"encoding/json"
	"fmt"

	"github.com/scottjbarr/pricebroadcaster/pkg/cache"
	"github.com/scottjbarr/pricebroadcaster/pkg/models"
	"github.com/scottjbarr/pricebroadcaster/pkg/publisher"
)

// Broadcaster is responsible for receiving and broadcasting
// data
type Broadcaster struct {
	Room      string
	Publisher publisher.Publisher
	Cache     cache.Cache
}

// New returns a new Broadcaster with the given Config
func New(room string, pub publisher.Publisher, cache cache.Cache) (*Broadcaster, error) {
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
