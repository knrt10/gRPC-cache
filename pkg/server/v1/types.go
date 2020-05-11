package v1

import (
	"sync"
	"time"
)

// Item store value of key and expiration
type Item struct {
	Object     interface{}
	Expiration int64
}

// Cache is a struct in which the cache's methods synchronize access to this map, so it is not
// recommended to keep any references to the map around after creating a cache.
type Cache struct {
	*cache
}

type worker struct {
	Interval time.Duration
	stop     chan bool
}

type cache struct {
	defaultExpiration time.Duration
	mu                sync.RWMutex
	store             map[interface{}]interface{}
	worker            *worker
}
