package util

import "github.com/go-redis/redis/v8"

var rdb *redis.Client

func NewRdb() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	return rdb
}
