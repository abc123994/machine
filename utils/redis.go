package utils

import (
	"context"
	"log"
	"machine_svc/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func RedisDB() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.Redis_conn,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.Redis_conn,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

	}
	return rdb

}
func RedisExample() {
	RedisDB().Set(context.Background(), "aaaa", "bbb", time.Second*100)
	res, err := RedisDB().Get(context.Background(), "aaaa").Result()
	if err == nil {
		log.Println(string(res))
	}
}
