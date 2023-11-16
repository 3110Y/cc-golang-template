#!/bin/sh
protoc --proto_path=api/proto --go_out=pkg/<SERVICE> --go_opt=paths=source_relative --go-grpc_out=pkg/<SERVICE> --go-grpc_opt=paths=source_relative api/proto/*.proto