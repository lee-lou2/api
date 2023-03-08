package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/lee-lou2/api/internal/errors"
	"os"
	"strconv"
	"time"
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

// GetValue Redis Get Value
func (c *NewClient) GetValue(key string) (string, error) {
	return c.Get(key).Result()
}

// SetValue Redis Set Value
func (c *NewClient) SetValue(key string, value string, expiration int) error {
	if err := c.Set(key, value, time.Duration(expiration)*time.Second).Err(); err != nil {
		return errors.New(errors.CacheDataConfigError, err)
	}
	return nil
}
