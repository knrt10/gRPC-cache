package v1

import (
	"runtime"
	"time"
)

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
func newCacheWithWorker(de time.Duration, ci time.Duration, m map[interface{}]interface{}) *Cache {
	c := newCache(de, m)
	C := &Cache{c}
	if ci > 0 {
		runWorker(c, ci)
		runtime.SetFinalizer(C, stopWorker)
	}
	return C
}

func newCache(de time.Duration, m map[interface{}]interface{}) *cache {
	c := &cache{
		defaultExpiration: de,
		store:             m,
	}
	return c
}

// This stops the Worker
func stopWorker(c *Cache) {
	c.worker.stop <- true
}

func runWorker(c *cache, ci time.Duration) {
	j := &worker{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.worker = j
	// This starts the new ticker and checks for expiration keys in cache
	go j.Run(c)
}

// Run deletes keys that get expired
func (j *worker) Run(c *cache) {
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			// This means key has expired
			c.deleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

// delete is used to delete key from the cache
func (c *cache) delete(k interface{}) (interface{}, bool) {
	delete(c.store, k)
	return nil, false
}

// deleteExpired is used to delete items from the cache when key time is expired
func (c *cache) deleteExpired() {
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.store {
		if v.(Item).Expiration > 0 && now > v.(Item).Expiration {
			c.delete(k)
		}
	}

	c.mu.Unlock()
}
