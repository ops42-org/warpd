# BINARY := warpd-builder warpd-deployer
BINARY := main
OSES := darwin linux
ARCHS := amd64
BINARYOSES := $(foreach o,$(OSES),$(BINARY:%=%-$(o)))
OSANDARCH := $(foreach a,$(ARCHS),$(BINARYOSES:%=%-$(a)))
BUILD ?= $(shell (git tag --points-at HEAD | tr '\n' ' ';git rev-parse HEAD | tr -d '\n'))
LDFLAGS = -ldflags '-X "github.com/ops42-org/warpd/cmd.version=$(BUILD)"'

.PHONY: build-all
build-all: build/warpd-darwin-amd64 build/warpd-linux-amd64

#build-all: $(OSANDARCH:%=build/%)

.PHONY: package
package: $(OSANDARCH:%=build/%.tar.gz)

#build/%.tar.gz: build/%
#	rm -f build/$(BINARY) && cp -v build/$(<:build/%=%) build/$(BINARY)
#	tar -cvzf $@ -C build $(BINARY)

build/warpd-darwin-amd64: main.go build
	GOOS=darwin ARCH=amd64 go build -o $@ $(LDFLAGS) $<

build/warpd-linux-amd64: main.go build
	GOOS=linux ARCH=amd64 go build -o $@ $(LDFLAGS) $<

clean:
	rm -rf ./build/*

build:
	mkdir -p ./build

.PHONY: test
test:
	go test -v ./...
