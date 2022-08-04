GO ?= go

# command to compiling the distributable. Specify GOOS and GOARCH for
# the target OS.
GO_DIST = GOOS=linux GOARCH=amd64 go build

.PHONY: dist build

all: clean build

dist:
	mkdir -p dist
	$(GO_DIST) -o dist/pricebroadcaster cmd/pricebroadcaster/main.go

build:
	mkdir -p build
	$(GO) build -o build/pricebroadcaster cmd/pricebroadcaster/main.go

run-example-publisher:
	go run cmd/price-publisher/main.go

run-example-http:
	go run cmd/price-http/main.go

test:
	$(GO) test ./...

install:
	$(GO) install ./...

clean:
	rm -rf build dist
