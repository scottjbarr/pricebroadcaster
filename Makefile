GO ?= go

# command to build and run on the local OS.
GO_BUILD = go build

# command to compiling the distributable. Specify GOOS and GOARCH for
# the target OS.
GO_DIST = GOOS=linux GOARCH=amd64 go build

.PHONY: dist

all: clean build

dist:
	mkdir -p dist
	$(GO_DIST) -o build/pricebroadcaster cmd/pricebroadcaster/main.go

build:
	mkdir -p build
	$(GO_BUILD) -o build/pricebroadcaster cmd/pricebroadcaster/main.go

run:
	go run cmd/pricebroadcaster/main.go

test:
	$(GO) test

clean:
	rm -rf build dist
