#!/bin/bash

terraform init terraform/

terraform apply -auto-approve terraform

HOST_INGRESS=$(terraform output ingress_host_ip)

ADDRESS_MAP_EXIST=$(grep "$HOST_INGRESS grpc-cache.example.com" /etc/hosts)

if [ "$ADDRESS_MAP_EXIST" == "" ]; then
  echo "Adding ingress host mapping to /etc/hosts"
  sudo -- sh -c -e "echo '$HOST_INGRESS grpc-cache.example.com' >> /etc/hosts";
fi

echo "Started all resources successfully, you can use the application now."
