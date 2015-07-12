# See http://peter.bourgon.org/go-in-production/
GO ?= go
CONFIG_FILE = ./conf/local.json
BIN = ./pricebroadcaster-go

all: clean build

build:
	$(GO) build

run: build
	$(BIN) -config $(CONFIG_FILE)

run-dev:
	$(GO) run main.go config.go redis.go -config $(CONFIG_FILE)

test:
	$(GO) test

clean:
	rm -f $(BIN)
