# node-exporter-plugin

build prometheus `node_exporter` to golang plugin

## usage

build plugin

```go
go build -buildmode=plugin -o prome.so node_exporter.go
```

start node_exporter with plugin mode

```go
go run plugin_loader.go --plugin=prome.so
```

## make

```bash
make
```
