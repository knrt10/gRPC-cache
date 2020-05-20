package server

import (
	"context"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/golang/protobuf/ptypes/empty"
	apis "github.com/knrt10/grpc-cache/proto"
)

const (
	bufSize = 1024 * 1024
	expire  = 10
	cleanup = 1
)

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	apis.RegisterCacheServiceServer(s, NewCacheService(time.Duration(expire)*time.Minute, time.Duration(cleanup)*time.Second))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAdd(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := apis.NewCacheServiceClient(conn)
	keyVal1 := &apis.Item{
		Key:        "kautilya",
		Value:      "knrt10",
		Expiration: "1m",
	}

	keyVal2 := &apis.Item{
		Key:        "2006",
		Value:      "percona",
		Expiration: "1m",
	}

	keyVal3 := &apis.Item{
		Key:        "foo",
		Value:      "bar",
		Expiration: "1m",
	}

	keyVal4 := &apis.Item{
		Key:        "temp",
		Value:      "bar",
		Expiration: "1µs",
	}

	c.Add(context.Background(), keyVal2)
	c.Add(context.Background(), keyVal3)
	c.Add(context.Background(), keyVal4)

	resp, err := c.Add(context.Background(), keyVal1)
	if err != nil {
		t.Fatalf("Adding key Failed: %v", err)
	}
	if resp.Key != "kautilya" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Key, "kautilya")
	}
	if resp.Value != "knrt10" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Key, "knrt10")
	}

	// Save keys
	// Checking for race condition
	for i := 0; i < 100; i++ {
		go c.Add(context.Background(), &apis.Item{
			Key:        strconv.Itoa(i),
			Value:      "Value of i is ",
			Expiration: strconv.Itoa(i),
		})
	}

}

func TestGet(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := apis.NewCacheServiceClient(conn)

	keyGet := &apis.GetKey{
		Key: "kautilya",
	}
	resp, err := c.Get(context.Background(), keyGet)
	if err != nil {
		t.Fatalf("Adding key Failed: %v", err)
	}
	if resp.Key != "kautilya" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Key, "kautilya")
	}
	if resp.Value != "knrt10" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Key, "knrt10")
	}
}

func TestGetAllItems(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := apis.NewCacheServiceClient(conn)

	_, err = c.GetAllItems(context.Background(), &empty.Empty{})
	if err != nil {
		t.Fatalf("Adding key Failed: %v", err)
	}
}

func TestDeleteKey(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := apis.NewCacheServiceClient(conn)

	keyGet := &apis.GetKey{
		Key: "22",
	}
	resp, err := c.DeleteKey(context.Background(), keyGet)
	if err != nil {
		t.Fatalf("Adding key Failed: %v", err)
	}
	if resp.Success != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Success, true)
	}
}

func TestDeleteAll(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := apis.NewCacheServiceClient(conn)

	resp, err := c.DeleteAll(context.Background(), &empty.Empty{})
	if err != nil {
		t.Fatalf("Adding key Failed: %v", err)
	}
	if resp.Success != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Success, true)
	}
}

// Testing deleted Key
func TestGetDeletedKey(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := apis.NewCacheServiceClient(conn)

	// Geting expired key
	keyGet := &apis.GetKey{
		Key: "temp",
	}
	_, err = c.Get(context.Background(), keyGet)
	if err.Error() != "rpc error: code = Unknown desc = No key found" {
		t.Errorf("Key not deleted")
	}
}

func TestDeleteKeyByExpiration(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := apis.NewCacheServiceClient(conn)
	keyVal1 := &apis.Item{
		Key:        "expired",
		Value:      "knrt10",
		Expiration: "1s",
	}

	resp, err := c.Add(context.Background(), keyVal1)
	if err != nil {
		t.Fatalf("Adding key Failed: %v", err)
	}
	if resp.Key != "expired" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Key, "expired")
	}
	if resp.Value != "knrt10" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Key, "knrt10")
	}

	time.Sleep(2 * time.Second)

	keyGet := &apis.GetKey{
		Key: "expired",
	}
	_, err = c.Get(context.Background(), keyGet)
	if err.Error() != "rpc error: code = Unknown desc = No key found" {
		t.Errorf("Key not deleted")
	}

}
