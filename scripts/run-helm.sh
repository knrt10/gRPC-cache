#!/bin/bash

# check namespace if present
{
  NAMESPACE_EXIST=$(kubectl get ns grpc-cache -o jsonpath='{.metadata.name}')
} &> /dev/null

if [ "$NAMESPACE_EXIST" != "grpc-cache" ]; then
  kubectl create ns grpc-cache
fi

# set current context namespace
kubectl config set-context $(kubectl config current-context) --namespace=grpc-cache

# create resources for k8s
helm install grpc-cache ./helm-charts

HOST_INGRESS=$(kubectl get ingress grpc-cache -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

# setup host
while true; do
  if [ "$HOST_INGRESS" != "" ]; then
    echo "Ingress address configured to: $HOST_INGRESS"
    break
  fi
  echo "Waiting for ingress to configure, sleeping for 10 seconds"
  HOST_INGRESS=$(kubectl get ingress grpc-cache -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
  sleep 10
done

ADDRESS_MAP_EXIST=$(grep "$HOST_INGRESS grpc-cache.example.com" /etc/hosts)

if [ "$ADDRESS_MAP_EXIST" == "" ]; then
  echo "Adding ingress host mapping to /etc/hosts"
  sudo -- sh -c -e "echo '$HOST_INGRESS grpc-cache.example.com' >> /etc/hosts";
fi

echo "Started all resources successfully, you can use the application now."
