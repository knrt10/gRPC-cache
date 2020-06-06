package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/knrt10/grpc-cache/proto"
)

var (
	address string
	conn    *grpc.ClientConn
	err     error
)

func main() {

	// Get address from flag
	flag.StringVar(&address, "addr", "127.0.0.1:5001", "Address on which you want to run server")
	flag.Parse()
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewCacheServiceClient(conn)

	// Add key
	keyVal1 := &api.Item{
		Key:        "22",
		Value:      "knrt10",
		Expiration: "-1m",
	}

	keyVal2 := &api.Item{
		Key:        "distributed",
		Value:      "systems",
		Expiration: "20s",
	}

	keyVal3 := &api.Item{
		Key:        "24",
		Value:      "Palash",
		Expiration: "2min10s",
	}

	keyVal4 := &api.Item{
		Key:        "prefixTest",
		Value:      "val1",
		Expiration: "10s",
	}

	keyVal5 := &api.Item{
		Key:        "prefixTest1",
		Value:      "val2",
		Expiration: "10s",
	}

	keyVal6 := &api.Item{
		Key:        "prefixTest2",
		Value:      "val3",
		Expiration: "10s",
	}

	c.Add(context.Background(), keyVal1)
	c.Add(context.Background(), keyVal2)
	c.Add(context.Background(), keyVal4)
	c.Add(context.Background(), keyVal5)
	c.Add(context.Background(), keyVal6)

	addKeyRes, err := c.Add(context.Background(), keyVal3)
	if err != nil {
		log.Fatalf("Error when calling Add: %s", err)
	}
	fmt.Println("Response from server for adding a key", addKeyRes)

	// Checking for race condition
	for i := 0; i < 50; i++ {
		go c.Add(context.Background(), &api.Item{
			Key:        strconv.Itoa(i),
			Value:      "Value of i is ",
			Expiration: strconv.Itoa(i),
		})
	}

	// Get key
	keyGet := &api.GetKey{
		Key: "distributed",
	}

	getKeyRes, err := c.Get(context.Background(), keyGet)
	if err != nil {
		log.Fatalf("Error when calling Get: %s", err)
	}
	fmt.Println("Response from server for getting a key", getKeyRes)

	// Get keys by prefix
	keyGetPrefix := &api.GetKey{
		Key: "prefixTest",
	}

	getKeyPrefixRes, err := c.GetByPrefix(context.Background(), keyGetPrefix)
	if err != nil {
		log.Fatalf("Error when calling Get: %s", err)
	}
	fmt.Println("Response from server for getting a keys by prefix", getKeyPrefixRes)

	// GetAllItems
	getAllKeysRes, err := c.GetAllItems(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("Error when calling GetAllItems: %s", err)
	}
	fmt.Println("Response from server for getting all keys", getAllKeysRes)

	// Delete Key
	deleteKeyRes, err := c.DeleteKey(context.Background(), keyGet)
	if err != nil {
		log.Fatalf("Error when calling DeleteKey: %s", err)
	}
	fmt.Println("Response from server after deleting a key", deleteKeyRes)

	// DeleteAll key
	deleteAllKeysResp, err := c.DeleteAll(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("Error when calling DeleteAll: %s", err)
	}
	fmt.Println("Response from server after deleting all keys", deleteAllKeysResp)

	// GetAllItems after deleting key
	_, err = c.GetAllItems(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Println("Response from server after no key found", err.Error())
	}
}
