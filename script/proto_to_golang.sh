#!/bin/bash

rm -rf ./generated/golang	

protoc ./proto/proto/user/**/*.proto --go_out=./generated/ --go-grpc_out=./generated/ --grpc-gateway_out ./generated/ 

protoc ./proto/proto/task/**/*.proto --go_out=./generated/ --go-grpc_out=./generated/ --grpc-gateway_out ./generated/
