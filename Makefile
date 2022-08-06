GO ?= go

# command to compiling the distributable. Specify GOOS and GOARCH for
# the target OS.
GO_DIST = GOOS=linux GOARCH=amd64 go build

.PHONY: dist build

all: clean build

dist:
	mkdir -p dist
	$(GO_DIST) -o dist/price-publisher cmd/price-publisher/main.go
	$(GO_DIST) -o dist/price-http cmd/price-http/main.go

build:
	mkdir -p build
	$(GO) build -o build/price-publisher cmd/price-publisher/main.go
	$(GO) build -o build/price-http cmd/price-http/main.go

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
