package service

import (
	"log"

	"github.com/scottjbarr/pricebroadcaster/pkg/broadcaster"
	"github.com/scottjbarr/pricebroadcaster/pkg/models"
)

type Service struct {
	Broadcaster *broadcaster.Broadcaster
}

func New(b *broadcaster.Broadcaster) *Service {
	return &Service{
		Broadcaster: b,
	}
}

func (s *Service) Run(ch chan *models.OHLC) error {
	for {
		select {
		case ohlc := <-ch:
			if err := s.Broadcaster.Broadcast(ohlc); err != nil {
				log.Printf("ERROR sending %+v", *ohlc)
			}
		}
	}

	return nil
}
