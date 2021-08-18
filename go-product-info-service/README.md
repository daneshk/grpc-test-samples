## To generate service stub file

```shell
protoc -I proto proto/product_info.proto  --go_out=plugins=grpc:ecommerce
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