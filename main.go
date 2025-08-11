package main

import (
	"fmt"
	"os"

	"github.com/ssd-81/RSS-feed-/internal/cli"
	"github.com/ssd-81/RSS-feed-/internal/config"
)

func main() {
	fmt.Println("start")
	temp := config.Read()
	state := config.State{}
	state.State = &temp
	cmds := cli.Commands{}
	// initialize map of handler functions for cmds
	cmds.Map := make(map[string]func(*config.State, cli.Command)) // ?????
	cmds.Register("login", cli.HandlerLogin)                      // ?????
	if len(os.Args) < 2 {
		fmt.Println("less than two command line arguments")
		os.Exit(0)
	}

}
