# node-exporter-plugin

build prometheus `node_exporter` to golang plugin

## usage

```
go build -buildmode=plugin -o prome.so node_exporter.go
go run plugin_loader.go prome.so
```
