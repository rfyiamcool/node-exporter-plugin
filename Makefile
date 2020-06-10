.PHONY: build start

all: build start

build:
	@go build -buildmode=plugin -o prome.so node_exporter.go

start:
	@go run plugin_loader.go --plugin=prome.so