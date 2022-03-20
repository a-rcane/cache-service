package main

import(
	"fmt"
	"github.com/go-redis/redis"
)


type redisDatabase struct {
	client *redis.Client 
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong,err)
}

func (r *redisDatabase) Set(key string, value string) (string, error) {
	_, err := r.client.Set(key, value, 0).Result()
	if err != nil {
		fmt.Println("Set is wrong")
	}
	return "sanchit"+key, nil
}

func (r *redisDatabase) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		fmt.Println("Get is wrong")
	}
	return value, nil
}
