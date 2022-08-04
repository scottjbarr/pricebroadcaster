# Price Broadcaster

Publish OHLC data to a Redis server.

This repository doesn't provide price sources. To use this repository you'll need to wire up data
sources.

This repo is useful to me because I can push prices to any number of services.

## Testing

```
$ go test
```

## Build

```
$ go build
```

## Running

Examples for how you could use this repository are provided.

See [the example config](conf/dev.env.example) if you want to run them.


```
$ source conf/your.env && go run cmd/price-publisher/main.go
```

A HTTP service to retrieve the latest price for a symbol.

There is only one endoint `GET /api/v1/prices/{symbol}`

```
$ source conf/your.env && go run cmd/price-http/main.go
```

## License

The MIT License (MIT)

Copyright (c) 2015-2022 Scott Barr

See [LICENSE](LICENSE)
