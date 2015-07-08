# See http://peter.bourgon.org/go-in-production/
GO ?= go
CONFIG_FILE = ./conf/local.json
BIN = pricebroadcaster-go

all: build

build:
	$(GO) build

run:
	$(GO) run main.go config.go redis.go -config $(CONFIG_FILE)

test:
	$(GO) test

clean:
	rm -f $(BIN)
