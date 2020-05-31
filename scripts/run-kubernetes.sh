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

POD_NAME=$(kubectl get po -n grpc-cache | cut -c -16 | tr "NAME" "\n")

kubectl port-forward -n grpc-cache $POD_NAME $HOST_PORT:$CONTAINER_PORT
