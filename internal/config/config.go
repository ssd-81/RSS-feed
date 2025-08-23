package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Db_url   string `json:"db_urL"`
	UserName string `json:"current_user_name"`
}


// type Command struct {
// 	Name      string
// 	Arguments []string
// }


func Read() *Config {

	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("home path unavailable")
		return &Config{}
	}
	path := filepath.Join(homePath, ".gatorconfig.json")
	config, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return &Config{}
	}
	var c Config
	err = json.Unmarshal(config, &c)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
		return &Config{}
	}
	return &c
}

func (c Config) SetUser(name string) {
	c.UserName = name
	newData, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		panic(err)
	}
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("home path unavailable")
		return
	}
	path := filepath.Join(homePath, ".gatorconfig.json")
	err = os.WriteFile(path, newData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")

}
