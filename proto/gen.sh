#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=../ *.proto
#install protoc and protoc-gen-go ~v1.28.1
#go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
#go install github.com/favadi/protoc-go-inject-tag@latest
cd ../internal/service
export PATH="$PATH:$(go env GOPATH)/bin"
protoc-go-inject-tag -input=./*.pb.go

#window: protoc-go-inject-tag -input="*.pb.go" -remove_tag_comment