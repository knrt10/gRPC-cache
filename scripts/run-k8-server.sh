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

# check namespace if present
{
  NAMESPACE_EXIST=$(kubectl get ns grpc-cache -o jsonpath='{.metadata.name}')
  KEY_EXIST=$(kubectl get secrets grpc-cache.example.com -o jsonpath='{.metadata.name}' -n grpc-cache)
  NGINX_CONTROLLER_EXIST=$(kubectl get svc nginx-ingress-grpc-cache-controller -o jsonpath='{.metadata.name}' -n grpc-cache)
  DEPLOYMENT_EXISTS=$(kubectl get deployment grpc-cache -n grpc-cache -o jsonpath='{.metadata.name}')
  SERVICE_EXISTS=$(kubectl get svc grpc-cache -n grpc-cache -o jsonpath='{.metadata.name}')
  INGRESS_EXISTS=$(kubectl get ingress grpc-cache -n grpc-cache -o jsonpath='{.metadata.name}')
} &> /dev/null

if [ "$NAMESPACE_EXIST" != "grpc-cache" ]; then
  kubectl create ns grpc-cache
fi

# set current context namespace
kubectl config set-context $(kubectl config current-context) --namespace=grpc-cache

if [ "$NGINX_CONTROLLER_EXIST" != "nginx-ingress-grpc-cache-controller" ]; then
  helm install stable/nginx-ingress --name-template=nginx-ingress-grpc-cache
fi

if [ "$KEY_EXIST" != "grpc-cache.example.com" ]; then
  openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=grpc-cache.example.com/O=grpc-cache.example.com"

  # create secret
  kubectl create secret tls grpc-cache.example.com --key tls.key --cert tls.crt -n grpc-cache
  rm tls.key tls.crt
fi

# create resources for k8s

if [ "$DEPLOYMENT_EXISTS" != "grpc-cache" ]; then
  kubectl create -f ./k8s/deployment.yaml
fi

if [ "$SERVICE_EXISTS" != "grpc-cache" ]; then
  kubectl create -f ./k8s/service.yaml
fi

if [ "$INGRESS_EXISTS" != "grpc-cache" ]; then
  kubectl create -f ./k8s/ingress.yaml
fi

