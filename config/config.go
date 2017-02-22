package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"errors"
)

type Config struct {
	Port string
	Greet string `json:"greet"`
	DB `json:"db"`
	RedisURL string `json:"redisurl"`
}

type DB struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Name string `json:"name"`
}

func (c *Config) Load() error {
	port := flag.String("port", "8080", "Server port")
	cFile := flag.String("conf", "config.json", "Config file")
	flag.Parse()

	raw, err := ioutil.ReadFile("config/" + *cFile)
	if err != nil {
		return err
	}

	json.Unmarshal(raw, &c)

	if c.Greet == "" || c.DB.Host == "" || c.RedisURL == "" {
		return errors.New("Loading config failed")
	}

	c.Port = *port

	return nil
}
