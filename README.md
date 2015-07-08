# Price Broadcaster

Poll and publish quote changes to a Redis server.

## Testing

    go test

## Build

    go build

## Install

    go get github.com/scottjbarr/pricebroadcaster-go

## Running

See `conf/example.json.sample` for an example config

    pricebroadcaster-go -c config/file.json

## License

The MIT License (MIT)

Copyright (c) 2015 Scott Barr

See [LICENSE.md](LICENSE.md)
