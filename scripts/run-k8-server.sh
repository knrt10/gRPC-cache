#!/bin/bash

# create resources for k8s
kubectl apply -f ./k8s/namespace.yaml
kubectl apply -f ./k8s/service.yaml
kubectl apply -f ./k8s/deployment.yaml
kubectl apply -f ./k8s/ingress.yaml

# set current context namespace
kubectl config set-context $(kubectl config current-context) --namespace=grpc-cache

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
