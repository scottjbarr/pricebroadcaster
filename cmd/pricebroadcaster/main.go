package main

import (
	"fmt"
	"os"

	"github.com/scottjbarr/pricebroadcaster"
)

// usage prints usage details
func usage() {
	fmt.Printf("Usage : %s symbol\n", os.Args[0])
}

func main() {
	config, err := pricebroadcaster.NewConfig()

	if err != nil {
		panic(err)
	}

	broadcaster, err := pricebroadcaster.New(*config)

	if err != nil {
		panic(err)
	}

	broadcaster.Start()
}
