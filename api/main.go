package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	cache "github.com/knrt10/grpc-cache/api/server"
	api "github.com/knrt10/grpc-cache/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	address string
	expire  int
	cleanup int
)

func main() {
	// Get address from flag
	flag.StringVar(&address, "addr", ":5001", "Address on which you want to run server")
	flag.IntVar(&expire, "exp", 10, "Default expiration duration of cache is 10 min")
	flag.IntVar(&cleanup, "cln", 5, "Cleanup interval duration of expired cache is 5 min")
	flag.Parse()

	opts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(200),
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer(opts...)
	// Default expiration of cache is 10 minutes and default purge time for expired items is 5 minutes
	api.RegisterCacheServiceServer(grpcServer, cache.NewCacheService(time.Duration(expire)*time.Minute, time.Duration(cleanup)*time.Minute))

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error in starting server %v", err)
	}
	fmt.Println("Started the server on:", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("err in serving gRPC %v\n", err)
	}
}
