package config

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "redis-18340.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:18340")
			if err != nil {
				log.Fatal("something went wrong in connection redis", err.Error())
			}

			return conn, err
		},
	}
}
