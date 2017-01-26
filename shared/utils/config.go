package util

import (
	"encoding/json"
	"fmt"
	"os"
)

//Configuration needed for plugins and bot
type Config struct {
	Token          string
	CommandPrefix  string
	BrainFile      string
	Administrators map[string]bool
}

func (c *Config) IsAdmin(user string) bool {
	_, ok := c.Administrators[user]
	return ok
}

func LoadConfig(f string) (Config, error) {

	file, err := os.Open(f)

	if err != nil {
		fmt.Println("Couldn't read config file")
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Couldn't parse json file")
		return Config{}, err
	}
	return config, err

}
