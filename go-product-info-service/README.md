## To generate service stub file

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative product_info.proto
```

## To build the Golang Package

```shell
go build -v -o bin/server
```

## To run the Go gRPC server

```shell
./bin/server
```

## To create docker image

```shell
docker build --tag go_grpc_simple_server .
```

## To run docker image

```shell
docker run -p 50051:50051 --rm -d --name go_grpc_simple_server -d go_grpc_simple_server
```

## To stop docker image

```shell
docker stop go_grpc_simple_server
```