package pricebroadcaster

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Redis config
type Redis struct {
	Host string
	Port int64
	DB   int64
	Room string
}

// Config top level configuration struct
type Config struct {
	SleepTime int64
	Redis
	Symbols []string
}

// NewConfig parses a config file
func NewConfig() (*Config, error) {
	c := Config{}

	c.Redis.Host = c.getEnv("REDIS_HOST", "127.0.0.1")
	c.Redis.Port = c.getEnvInt("REDIS_PORT", 6379)
	c.Redis.DB = c.getEnvInt("REDIS_DB", 0)

	c.Redis.Room = os.Getenv("ROOM")

	if c.Redis.Room == "" {
		return nil, fmt.Errorf("ROOM must be supplied")
	}

	c.SleepTime = c.getEnvInt("SLEEP_TIME", 1000)
	c.Symbols = strings.Split(os.Getenv("SYMBOLS"), " ")

	if len(c.Symbols) == 0 || (len(c.Symbols) == 1 && c.Symbols[0] == "") {
		return nil, fmt.Errorf("SYMBOLS must be supplied")
	}

	return &c, nil
}

func (c Config) getEnvInt(name string, value int64) int64 {
	s := os.Getenv(name)

	if s == "" {
		return value
	}

	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		log.Printf("WARN %v is not an integer. Using default %v", name, value)
		return value
	}

	return i
}

func (c Config) getEnv(name, value string) string {
	s := os.Getenv(name)

	if s == "" {
		return value
	}

	return s
}
