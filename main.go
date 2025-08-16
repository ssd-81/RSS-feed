package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/ssd-81/RSS-feed-/internal/cli"
	"github.com/ssd-81/RSS-feed-/internal/config"
)

func main() {
	temp := config.Read()
	state := config.State{}
	state.State = &temp
	cmds := cli.Commands{}
	// initialize map of handler functions for cmds
	cmds.Map = make(map[string]func(*config.State, cli.Command) error) // ?????
	cmds.Register("login", cli.HandlerLogin)                           // ?????
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("error: not enough arguments were provided")
		os.Exit(1)
	} else if len(argsWithoutProg) == 1 {
		fmt.Println("error: username is required")
		os.Exit(1)
	} else if len(argsWithoutProg) == 2 {
		temp.SetUser(argsWithoutProg[1])
	}

	fmt.Println(config.Read())

}
