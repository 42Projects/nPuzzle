#!/bin/sh
protoc -I proto proto/npuzzle.proto --go_out=plugins=grpc:proto ; protoc proto/npuzzle.proto --proto_path=proto --js_out=import_style=commonjs:client/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:client/src/
