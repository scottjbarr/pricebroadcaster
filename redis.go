package pricebroadcaster

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// Return a formatted Redis connection string
func (config *Redis) connString() string {
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}

// Return a connection to a Redis server
//
// The caller will need to manage closing the connection.
func connect(config Config) (redis.Conn, error) {
	c, err := redis.Dial("tcp", config.Redis.connString())

	if err != nil {
		return c, err
	}

	// switch to the correct database
	if _, err = c.Do("SELECT", config.Redis.DB); err != nil {
		return c, nil
	}

	return c, nil
}
