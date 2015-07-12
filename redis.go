package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Return a formatted Redis connection string
func (config *Redis) ConnString() string {
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}

// Return a connection to a Redis server
//
// The caller will need to manage closing the connection.
func Connect(config *Redis) (redis.Conn, error) {
	c, err := redis.Dial("tcp", config.ConnString())

	if err != nil {
		return c, err
	}

	// switch to the correct database
	if _, err = c.Do("SELECT", config.DB); err != nil {
		return c, nil
	}

	return c, nil
}
