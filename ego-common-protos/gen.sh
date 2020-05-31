#!/usr/bin/env bash

# generate the gRPC code
#protoc -I. -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --go_out=plugins=grpc,paths=source_relative:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.  ./user/user_info/user_info.proto

#user
protoc -I.  --go_out=plugins=grpc:.   ./modules/user/user_info/user_info.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/user/user_addr/user_addr.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/user/user_balance_log/user_balance_log.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/user/user_level/user_level.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/user/user_login_log/user_login_log.proto

#inventory
protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/brand/brand.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product/product.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product_attribute/product_attribute.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product_attribute_category/product_attribute_category.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product_attribute_value/product_attribute_value.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product_category/product_category.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product_full_reduction/product_full_reduction.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product_sku_stock/product_sku_stock.proto

protoc -I.  --go_out=plugins=grpc:.   ./modules/inventory/product_vertify_record/product_vertify_record.proto
