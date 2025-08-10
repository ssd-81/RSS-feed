package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Db_url   string `json:"db_urL"`
	UserName string `json:"current_user_name"`
}

type State struct {
	State *Config
}

func Read() Config {

	config, err := os.ReadFile(".gatorconfig.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return Config{}
	}
	var c Config
	err = json.Unmarshal(config, &c)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
		return Config{}
	}
	return c
}

func (c Config) SetUser(name string) {
	c.UserName = name
	newData, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(".gatorconfig.json", newData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")

}
