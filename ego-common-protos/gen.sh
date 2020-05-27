#!/usr/bin/env bash

# generate the gRPC code
#protoc -I. -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --go_out=plugins=grpc,paths=source_relative:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.  ./user/user_info/user_info.proto

protoc -I. -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:.  ./user/user_info/user_info.proto