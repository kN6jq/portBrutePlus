package plugins

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func ScanRedis(ip, port, username, password string) (err error, result bool) {
	result = false
	rdb := redis.NewClient(&redis.Options{
		Addr:     ip + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	pong := rdb.Ping(context.Background())
	if pong.Err() == nil {
		result = true
	}

	return pong.Err(), result
}
