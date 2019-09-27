# birpc

[![Go Report Card](https://goreportcard.com/badge/github.com/gocs/birpc)](https://goreportcard.com/report/github.com/gocs/birpc)

## goal

simple grpc golang game development in ebiten\
lol just capturing cursor

## pre-requisite

`go >1.12`

`protoc` from [protocolbuffers/protobuf releases binary](https://github.com/protocolbuffers/protobuf/releases)

save it to `$GOPATH/bin` or `%GOPATH\bin%`

## running

```
protoc --proto_path=src/proto --go_out=plugins=grpc:src/proto mouse.proto
```
```
go get
go run ./src/server/server.go
```
another terminal
```
go run ./src/client/client.go
```

## status

needs more knowledge about game design

needs more knowledge about concurrency in go

barely sync\
esp. client receive

needs update of collision

## LICENSE

apache 2.0
