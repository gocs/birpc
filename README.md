# birpc

[![Go Report Card](https://goreportcard.com/badge/github.com/gocs/birpc)](https://goreportcard.com/report/github.com/gocs/birpc)
[![Go 1.12](https://img.shields.io/badge/go-1.12-7cf.svg)](https://golang.org/dl/)
[![Build Status](https://travis-ci.org/gocs/birpc.svg?branch=master)](https://travis-ci.org/gocs/birpc)

## goal

simple grpc golang game development in ebiten\
lol just capturing cursor

## pre-requisite

`go >1.12`

`protoc` from [protocolbuffers/protobuf releases binary](https://github.com/protocolbuffers/protobuf/releases)

save it to `$GOPATH/bin` in mac or `%GOPATH%\bin` in windows

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
