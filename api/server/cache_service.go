package server

import (
	"runtime"
	"sync"
	"time"
)

// Item items value of key and expiration
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
	items             map[interface{}]interface{}
	worker            *worker
}

// NewCacheService is used to initialize a new cache.
// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one,
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.deleteExpired().
func NewCacheService(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[interface{}]interface{})
	return newCacheWithWorker(defaultExpiration, cleanupInterval, items)
}

// This ensures Worker goroutine (which granted it
// was enabled is running deleteExpired on c forever) does not keep
// the returned C object from being garbage collected. When it is
// garbage collected, the finalizer stops the Worker goroutine, after
// which c can be collected.
func newCacheWithWorker(defaultExpiration time.Duration, cleanupInterval time.Duration, items map[interface{}]interface{}) *Cache {
	c := newCache(defaultExpiration, items)
	C := &Cache{c}
	if cleanupInterval > 0 {
		runWorker(c, cleanupInterval)
		runtime.SetFinalizer(C, stopWorker)
	}
	return C
}

func newCache(defaultExpiration time.Duration, items map[interface{}]interface{}) *cache {
	c := &cache{
		defaultExpiration: defaultExpiration,
		items:             items,
	}
	return c
}

// This stops the Worker
func stopWorker(c *Cache) {
	c.worker.stop <- true
}

func runWorker(c *cache, cleanupInterval time.Duration) {
	w := &worker{
		Interval: cleanupInterval,
		stop:     make(chan bool),
	}
	c.worker = w
	// This starts the new ticker and checks for expiration keys in cache
	go w.Run(c)
}

// Run deletes keys that get expired
func (w *worker) Run(c *cache) {
	ticker := time.NewTicker(w.Interval)
	for {
		select {
		case <-ticker.C:
			// This means key has expired
			c.deleteExpired()
		case <-w.stop:
			ticker.Stop()
			return
		}
	}
}

// delete is used to delete key from the cache
func (c *cache) delete(k interface{}) (interface{}, bool) {
	delete(c.items, k)
	return nil, false
}

// deleteExpired is used to delete items from the cache when key time is expired
func (c *cache) deleteExpired() {
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		if v.(Item).Expiration > 0 && now > v.(Item).Expiration {
			c.delete(k)
		}
	}

	c.mu.Unlock()
}
