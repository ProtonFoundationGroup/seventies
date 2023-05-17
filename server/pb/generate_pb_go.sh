#!/bin/zsh

protoc -I=. --gogofaster_out=:../pb-go --gogofaster_opt=paths=source_relative api/*.proto

protoc -I=.:${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
                --gogofaster_out=plugins=grpc:../pb-go \
                --gogofaster_opt=paths=source_relative \
                --grpc-gateway_opt=paths=source_relative \
                --grpc-gateway_out=logtostderr=true:../pb-go \
                --swagger_out=logtostderr=true:../pb-go \
                api/*.proto