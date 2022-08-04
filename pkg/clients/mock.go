package clients

import (
	"time"

	"github.com/scottjbarr/pricebroadcaster/pkg/models"
)

type MockClient struct {
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (m *MockClient) Get(symbol string) (*models.OHLC, error) {
	ohlc := models.OHLC{
		Symbol: symbol,
		Close:  1234.56,
		Time:   time.Now().Unix(),
	}

	return &ohlc, nil
}
