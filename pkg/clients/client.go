package clients

import (
	"github.com/scottjbarr/pricebroadcaster/pkg/models"
)

type Client interface {
	Get(string) (*models.OHLC, error)
}
