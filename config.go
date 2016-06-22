package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Redis struct {
	Host string
	Port int
	DB   int
	Room string
}

// Config top level configuration struct
type Config struct {
	SleepTime int `yaml:"sleep_time"`
	Redis
	Symbols []string
}

// ParseConfig parses a config file
func ParseConfig(filename string) *Config {
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		panic(err)
	}

	return &config
}
