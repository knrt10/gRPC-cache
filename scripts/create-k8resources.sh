#!/bin/bash

# get variables for ports
HOST_PORT=$1
CONTAINER_PORT=$2

if [ "$HOST_PORT" == "" ]; then
  HOST_PORT="5001"
fi

if [ "$CONTAINER_PORT" == "" ]; then
  CONTAINER_PORT="5001"
fi

# create namespace for application
kubectl create ns grpc-cache
kubectl run grpc-cache --image=knrt10/grpc-cache --port=$CONTAINER_PORT --generator=run/v1 -n grpc-cache
kubectl expose rc grpc-cache --type=LoadBalancer --name=grpc-cache -n grpc-cache
