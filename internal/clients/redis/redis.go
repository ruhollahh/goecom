package redisclient

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func New(config Config) *redis.Pool {
	fmt.Println(fmt.Sprint(config.Host, ":", config.Port))
	client := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprint(config.Host, ":", config.Port),
				redis.DialDatabase(config.DB),
				redis.DialPassword(config.Password),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return client
}
