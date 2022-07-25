# WISHLIST GRPC MICROSERVICE

## Overview

GRPC microservice to manage a wishlist
The microservice has been build following the DDD design pattern.

protoc --proto_path=./proto --go_out=paths=source_relative:pb/ --go-grpc_out=paths=source_relative:pb/ proto\*.proto

### How to run

```sh
go run cmd\server\main.go
go run cmd\client\main.go
```

### Test

```sh
go test -v .\test\
```

### Docker

```sh
docker-compose up --build
```
