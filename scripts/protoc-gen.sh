#!/bin/bash

if [ ! -d "pkg/api/v1" ]
then
  mkdir -p pkg/api/v1
fi
protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:pkg/api/v1 cache-service.proto
