package broadcaster

import (
	"encoding/json"

	"github.com/scottjbarr/pricebroadcaster/pkg/cache"
	"github.com/scottjbarr/pricebroadcaster/pkg/models"
	"github.com/scottjbarr/pricebroadcaster/pkg/publisher"
)

// Broadcaster is responsible for receiving and broadcasting data
type Broadcaster struct {
	Publisher publisher.Publisher
	Cache     cache.Cache
}

// New returns a new Broadcaster.
func New(pub publisher.Publisher) *Broadcaster {
	return NewWithCache(pub, cache.NewNoopCache())
}

// New returns a new Broadcaster with a cache.
func NewWithCache(pub publisher.Publisher, cache cache.Cache) *Broadcaster {
	return &Broadcaster{
		Publisher: pub,
		Cache:     cache,
	}
}

// Broadcast OHLC data.
func (b *Broadcaster) Broadcast(ohlc *models.OHLC) error {
	// format the JSON with a root element
	data, err := json.Marshal(ohlc)
	if err != nil {
		return err
	}

	if err := b.Publisher.Publish(data); err != nil {
		return err
	}

	return b.Cache.Set(ohlc.Symbol, data)
}
