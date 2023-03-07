package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

type NewClient struct {
	*redis.Client
}

// RedisClient Redis 클라이언트
func RedisClient() *NewClient {
	dsn := fmt.Sprintf(
		"%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))
	client := &NewClient{redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})}

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}

// Pub Redis Publish
func (c *NewClient) Pub(channel string, message interface{}) error {
	err := c.Publish(channel, message).Err()
	if err != nil {
		return err
	}

	return nil
}

// Sub Redis Subscribe
func (c *NewClient) Sub(channel string, message chan<- string) error {
	sub := c.Subscribe(channel)

	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage()
		if err != nil {
			return err
		}

		message <- msg.Payload
	}
}

// Produce Redis Producer
func (c *NewClient) Produce(key string, value interface{}) error {
	_, err := c.LPush(key, value).Result()
	if err != nil {
		return err
	}

	return nil
}

// Consume Redis Consumer
func (c *NewClient) Consume(key string) (string, error) {
	result, err := c.BRPop(0, key).Result()
	if err != nil {
		return "", err
	}

	return result[1], nil
}
