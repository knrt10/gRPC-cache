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
	flag.StringVar(&address, "addr", ":5001", "Address on which you want to run server")
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
		Key:        "2006",
		Value:      "percona",
		Expiration: "20s",
	}

	keyVal3 := &api.Item{
		Key:        "24",
		Value:      "Palash",
		Expiration: "2min10s",
	}

	c.Add(context.Background(), keyVal1)
	c.Add(context.Background(), keyVal2)

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
		Key: "2006",
	}

	getKeyRes, err := c.Get(context.Background(), keyGet)
	if err != nil {
		log.Fatalf("Error when calling Get: %s", err)
	}
	fmt.Println("Response from server for getting a key", getKeyRes)

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
