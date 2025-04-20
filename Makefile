BINARY_ARCHS ?= amd64 arm64 wasm
MACH ?= $(shell uname -m)

# build all architectures
build-all: $(foreach arch,$(BINARY_ARCHS),build-$(arch))

# build for the current architecture
build: build-$(MACH)

# build for the specified architecture
build-amd64:
	GOARCH=amd64 GOOS=linux go build -o bin/$(MACH)/selfctl .
build-arm64:
	GOARCH=arm64 GOOS=linux go build -o bin/$(MACH)/selfctl .
build-wasm:
	GOOS=js GOARCH=wasm go build -o bin/$(MACH)/selfctl.wasm .
build-windows:
	GOARCH=amd64 GOOS=windows go build -o bin/$(MACH)/selfctl.exe .
