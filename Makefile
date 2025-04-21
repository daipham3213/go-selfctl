BINARY_ARCHS ?= amd64 arm64 wasm windows i386
MACH ?= $(shell uname -m)

# build all architectures
build-all: $(foreach arch,$(BINARY_ARCHS),build-$(arch))

# build for the current architecture
build: build-$(MACH)

# build for the specified architecture
build-amd64:
	GOARCH=amd64 GOOS=linux go build -o bin/amd64/selfctl-amd64 main.go
build-arm64:
	GOARCH=arm64 GOOS=linux go build -o bin/arm64/selfctl-arm64 main.go
build-i386:
	GOARCH=386 GOOS=linux go build -o bin/i386/selfctl-i386 main.go
build-wasm:
	GOOS=js GOARCH=wasm go build -o bin/wasm/selfctl.wasm main.go
build-windows:
	GOARCH=amd64 GOOS=windows go build -o bin/windows/selfctl.exe main.go
