package dbops

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"log"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

type Value struct {
	Id       int
	Username string
	Email    string
}

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("数据库 redis 连接成功")
}

func SetValue(key string, value interface{}) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetValue(key string) ([]byte, error) {
	val, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return val, nil
}
