# node-exporter-plugin

prome node exporter build plugin

## usage

```
go build -buildmode=plugin -o prome.so node_exporter.go
go run plugin_loader.go prome.so
```
