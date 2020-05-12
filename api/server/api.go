package server

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/knrt10/percona-cache/proto"
)

// Add is used to add key/value pair to the cache.
func (c *cache) Add(ctx context.Context, item *api.Item) (*api.Item, error) {
	var expiration int64
	duration, _ := time.ParseDuration(item.Expiration)
	// Meaning d is of form "2m30s"
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}
	c.mu.Lock()
	c.items[item.Key] = Item{
		Object:     item.Value,
		Expiration: expiration,
	}
	c.mu.Unlock()
	return item, nil
}

// Get method is used to key/value pair while providing key as args
func (c *cache) Get(ctx context.Context, args *api.GetKey) (*api.Item, error) {
	key := args.Key
	// Locking so that other goroutines cannot access this at the same time
	c.mu.RLock()
	value, exists := c.items[key]
	// No key found
	if !exists {
		c.mu.RUnlock()
		return nil, ErrNoKey
	}

	// This means key has some expiration
	if value.(Item).Expiration > 0 {
		if time.Now().UnixNano() > value.(Item).Expiration {
			c.mu.RUnlock()
			return nil, ErrKeyExpired
		}
	}
	c.mu.RUnlock()
	return &api.Item{
		Key:        key,
		Value:      value.(Item).Object.(string),
		Expiration: time.Unix(0, value.(Item).Expiration).String(),
	}, nil
}

// GetAllItems method get all unexpired keys from the cache
func (c *cache) GetAllItems(ctx context.Context, in *empty.Empty) (*api.AllItems, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var items []*api.Item
	now := time.Now().UnixNano()
	for k, v := range c.items {
		if v.(Item).Expiration > 0 {
			if now > v.(Item).Expiration {
				continue
			}
		}

		items = append(items, &api.Item{
			Key:        k.(string),
			Value:      v.(Item).Object.(string),
			Expiration: time.Unix(0, v.(Item).Expiration).String(),
		})
	}

	// This means no keys were found, or all were expired
	if len(items) < 1 {
		return nil, ErrNoKey
	}

	return &api.AllItems{
		Items: items,
	}, nil
}

// DeleteKey deletes an item from the cache. Does nothing if the key is not in the cache.
func (c *cache) DeleteKey(ctx context.Context, args *api.GetKey) (*api.Success, error) {
	c.mu.Lock()
	c.delete(args.Key)
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

// Delete all items from the cache.
func (c *cache) DeleteAll(ctx context.Context, in *empty.Empty) (*api.Success, error) {
	c.mu.Lock()
	c.items = map[interface{}]interface{}{}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}
