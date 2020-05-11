<p align="center">
  <img src="https://user-images.githubusercontent.com/24803604/81601158-2de68800-93e8-11ea-838c-901f44245498.png" />
</p>

> In memory cache, using gRPC

[![Build Status](https://travis-ci.org/knrt10/gRPC-cache.svg?branch=master)](https://travis-ci.org/knrt10/gRPC-cache)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/knrt10/gRPC-cache/pkg/server/v1)

## Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Usage](#usage)
    - [Local](#local)
    - [Docker](#docker)
- [API](#api)
    - [Add](#add)
    - [Get](#get)
    - [GetAllItems](#getallitems)
    - [DeleteKey](#deletekey)
    - [DeleteAll](#deleteall)
- [Testing](#testing)

## Requirements


- [Golang min-version(1.11)](https://golang.org/)
- make
- [protobuf](https://github.com/golang/protobuf)


## Getting Started

Once you've cloned this repo, run these commands in this directory:

```bash
# Only needed the first time:
$ make all

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

You can run server in 2 ways

### Local

- Start your server `./server-cache` or `make server`
- `./server-cache -addr=":5001"` to run server on port `5001`

### Docker

- Run this command `make docker` // This will build the image
- Run the server using `make dockerServer`

After running the server, start your client `./client-cache` or `make client` in a different terminal

- `./client-cache -addr=":5001"` is server is running running on port `5001`

## API

Proto syntax `proto3` is used. You can find the [proto file here](https://github.com/knrt10/gRPC-cache/tree/master/api/proto/v1/cache-service.proto)

### Add

This is used to add key/value to the cache

**Params**:- 

```go
Key = string
Value = string
Expiration = string

// Example
&api.Item{
	Key:        "22",
	value:      "knrt10",
	Expiration: "1m",
}
```

### Get

This is used to get key value pair for a particular key

**Params**:- 

```go
Key = string
// Example
&api.GetKey{
	Key: "23",
}
```

### GetAllItems

Used to get all key value pairs


### DeleteKey

Used to delete item by a particular key from the cache

**Params**:- 

```go
Key = string
// Example
&api.GetKey{
	Key: "23",
}
```

### DeleteAll

Used to clear the whole cache

## Testing

After running `make all` just run `make test` to run the tests. It has **coverage of 80.9%**

```bash
go test pkg/server/v1/* -v -cover
=== RUN   TestAdd
--- PASS: TestAdd (3.97s)
=== RUN   TestGet
--- PASS: TestGet (0.00s)
=== RUN   TestGetAllItems
--- PASS: TestGetAllItems (0.15s)
=== RUN   TestDeleteKey
--- PASS: TestDeleteKey (0.00s)
=== RUN   TestDeleteAll
--- PASS: TestDeleteAll (0.01s)
=== RUN   TestGetDeletedKey
--- PASS: TestGetDeletedKey (0.00s)
PASS
coverage: 80.9% of statements
ok  	command-line-arguments	5.289s	coverage: 80.9% of statements
```
