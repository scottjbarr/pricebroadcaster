package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

// Return a formatted Redis connection string
func (config *Redis) ConnString() string {
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}

// Return a connection to a Redis server
func Connect(config *Redis) redis.Conn {
	c, err := redis.Dial("tcp", config.ConnString())

	if err != nil {
		log.Fatal(err)
	}

	// switch to the correct database
	if _, err = c.Do("SELECT", config.DB); err != nil {
		log.Fatal(err)
	}

	return c
}
