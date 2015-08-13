package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-telegram-bot/yatzie/shared/registry"
)

func LoadConfig(f string) (plugin_registry.Config, error) {

	file, err := os.Open(f)

	if err != nil {
		fmt.Println("Couldn't read config file")
		return plugin_registry.Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config plugin_registry.Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Couldn't parse json file")
		return plugin_registry.Config{}, err
	}
	return config, err

}
