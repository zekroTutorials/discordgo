package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func ParseConfigFromJSONFile(fileName string) (c *Config, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	c = new(Config)
	err = json.NewDecoder(f).Decode(c)

	return
}
