#!/bin/bash

protoc demo.proto --go_out=../server --go-grpc_out=../server
protoc demo.proto --go_out=../client --go-grpc_out=../client


protoc repeated.proto --go_out=../server --go-grpc_out=../server
protoc repeated.proto --go_out=../client --go-grpc_out=../client

# protoc --proto_path=server --go_out=plugins=grpc:. proto/demo.proto
# protoc subscription.proto --go_out=../server --go-grpc_out=../server