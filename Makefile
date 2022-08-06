GO ?= go

# command to compiling the distributable. Specify GOOS and GOARCH for
# the target OS.
GO_DIST = GOOS=linux GOARCH=amd64 go build

.PHONY: dist build

all: clean build dist

prepare:
	mkdir -p build dist

dist: prepare dist-broadcaster dist-http

dist-broadcaster:
	$(GO_DIST) -o dist/price-broadcaster cmd/price-broadcaster/main.go

dist-http:
	$(GO_DIST) -o dist/price-http cmd/price-http/main.go

build: prepare build-broadcaster build-http

build-broadcaster:
	$(GO) build -o build/price-broadcaster cmd/price-broadcaster/main.go

build-http:
	$(GO) build -o build/price-http cmd/price-http/main.go

run-example-broadcaster:
	go run cmd/price-broadcaster/main.go

run-example-http:
	go run cmd/price-http/main.go

test:
	$(GO) test ./...

install:
	$(GO) install ./...

clean:
	rm -rf build dist
