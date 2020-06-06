<p align="center">
  <img src="https://user-images.githubusercontent.com/24803604/81805297-3948c900-9538-11ea-82d0-38a4aee7eb10.png" />
</p>

> In memory cache, using gRPC

[![Build Status](https://travis-ci.org/knrt10/gRPC-cache.svg?branch=master)](https://travis-ci.org/knrt10/gRPC-cache)
[![Coverage Status](https://coveralls.io/repos/github/knrt10/gRPC-cache/badge.svg)](https://coveralls.io/github/knrt10/gRPC-cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/knrt10/grpc-Cache)](https://goreportcard.com/report/github.com/knrt10/grpc-Cache)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/knrt10/gRPC-cache/api/server)

## Contents

- [About](#about)
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Usage](#usage)
    - [Local](#local)
    - [Docker](#docker)
    - [Kubernetes](#kubernetes)
- [API](#api)
    - [Add](#add)
    - [Get](#get)
    - [GetByPrefix](#getByPrefix)
    - [GetAllItems](#getallitems)
    - [DeleteKey](#deletekey)
    - [DeleteAll](#deleteall)
- [Testing](#testing)
- [Example](#example)

## About

Go in memory cache using gRPC to generate API. Functionalities include

- Adding/Replacing key/value
- Getting value using a key
- Getting all keys
- Deleting a particular key
- Deleting all keys
- **Concurrency safe** and on client side calls can be made using goroutines

## Requirements

- [Golang min-version(1.11)](https://golang.org/)
- make
- [protobuf](https://github.com/golang/protobuf)
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) installed locally
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) installed locally
- [helm](https://helm.sh/) installed locally
- [grpcurl](https://github.com/fullstorydev/grpcurl) for testing server running on k8s

## Getting Started

Once you've cloned this repo, run these commands in this directory:

```bash
# Only needed the first time:
$ make build

# Then run to start server
$ ./server-cache --help

Usage of ./server-cache:
  -addr string
      Address on which you want to run server (default ":5001")
  -cln int
      Cleanup interval duration of expired cache is 5 min (default 5)
  -exp int
      Default expiration duration of cache is 10 min (default 10)

# To use client
$ ./client-cache --help

Usage of ./client-cache:
  -addr string
      Address on which you want to run server (default ":5001")
```

## Usage

You can run server in 2 ways either using local setup or docker.

### Local

- Start your server `./server-cache` or `make server`
- `./server-cache -addr=":5001"` to run server on port `5001`

### Docker

- Run this command `make docker` // This will build the image
- Run the server using `make dockerServer`

After running the server, start your client `./client-cache` or `make client` in a different terminal

- `./client-cache -addr=":5001"` is server is running running on port `5001`

### Kubernetes

You can run server on your kubernetes cluster. All the resources are created on `grpc-cache` namespace. Before running the command below make sure your cluster is up and running.

- Run this command `make run-k8s-server`

This will create all the required resources needed to run your grpc server. Make sure all resources are ready before running the below command to get your IP for ingress.

`kubectl get ingress grpc-cache -o jsonpath='{.status.loadBalancer.ingress[0].ip}'`
> 192.168.64.3

You will have a diffrent output. You need this to add to your `/etc/hosts`. Add the above IP to your `/etc/hosts` to map to your ingress host.

`192.168.64.3 grpc-cache.example.com`
> 192.168.64.3 will be different for you.

Now you can easily test it using [grpcurl](https://github.com/fullstorydev/grpcurl).

#### Example

```bash
# To list all services exposed by a server, use the "list" verb.
grpcurl --insecure grpc-cache.example.com:443 list

# The "describe" verb will print the type of any symbol that the server knows about or that is found in a given protoset file. It also prints a description of that symbol, in the form of snippets of proto source. It won't necessarily be the original source that defined the element, but it will be equivalent.

grpcurl --insecure grpc-cache.example.com:443 list cacheService.CacheService

# To add a key
grpcurl --insecure -d '{"key": "knrt10", "value": "pro", "expiration": "3m"}' grpc-cache.example.com:443 cacheService.CacheService/Add

# To get key
grpcurl --insecure -d '{"key": "knrt10"}' grpc-cache.example.com:443 cacheService.CacheService/Get

# To get all keys
grpcurl --insecure grpc-cache.example.com:443 cacheService.CacheService/GetAllItems
```

Similarly you can use all the methods as shown in API below

## API

Proto syntax `proto3` is used. You can find the [proto file here](https://github.com/knrt10/gRPC-cache/tree/master/proto/cache-service.proto)

### Add

This is used to add key/value to the cache

```go
func (c Cache) Add(ctx context.Context, item *api.Item) (*api.Item, error)
```

### Get

This is used to get key value pair for a particular key

```go
func (c Cache) Get(ctx context.Context, args *api.GetKey) (*api.Item, error)
```

### GetByPrefix

Used to get all key value pairs by prefix

```go
func (c Cache) GetByPrefix(ctx context.Context, args *api.GetKey) (*api.AllItems, error)
```

### GetAllItems

Used to get all key value pairs

```go
func (c Cache) GetAllItems(ctx context.Context, in *empty.Empty) (*api.AllItems, error)
```

### DeleteKey

Used to delete item by a particular key from the cache

```go
func (c Cache) DeleteKey(ctx context.Context, args *api.GetKey) (*api.Success, error)
```

### DeleteAll

Used to clear the whole cache

```go
func (c Cache) DeleteAll(ctx context.Context, in *empty.Empty) (*api.Success, error)
```

## Testing

After running `make build` just run `make test` to run the tests. It has **coverage of 92.7%**

```bash
go test api/server/* -v -cover -race
=== RUN   TestAdd
--- PASS: TestAdd (0.03s)
=== RUN   TestGet
--- PASS: TestGet (0.01s)
=== RUN   TestGetByPrefix
--- PASS: TestGetByPrefix (0.01s)
=== RUN   TestGetAllItems
--- PASS: TestGetAllItems (0.01s)
=== RUN   TestDeleteKey
--- PASS: TestDeleteKey (0.00s)
=== RUN   TestDeleteAll
--- PASS: TestDeleteAll (0.00s)
=== RUN   TestGetDeletedKey
--- PASS: TestGetDeletedKey (0.01s)
=== RUN   TestDeleteKeyByExpiration
--- PASS: TestDeleteKeyByExpiration (2.01s)
PASS
coverage: 92.7% of statements
ok  	command-line-arguments	3.709s	coverage: 92.7% of statements
```

## Example

Please refer to [examples](https://github.com/knrt10/gRPC-cache/tree/master/examples) directory for more information
