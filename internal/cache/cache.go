package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration)
}

var client Cache = nil

// createClient creates the default redis client
func createClient() {
	rc := &redisCache{}
	rc.start()
	client = rc
}

// GetClient returns singleton instance of cache client;
// always use this for cache get/set.
func GetClient() Cache {
	if client == nil {
		createClient()
	}
	return client
}

// func CacheUsage() {
// 	c := cache.GetClient()
// 	if c == nil {
// 		fmt.Println("empty client returned")
// 		return
// 	}
// 	c.Set(context.Background(), "some_key", "random data", time.Hour)
// 	val, err := c.Get(context.Background(), "some_key")
// 	if err != nil {
// 		fmt.Println("error came: " + err.Error())
// 		return
// 	}
// 	fmt.Printf("returned val: %+v", val)
// }
