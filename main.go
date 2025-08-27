package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/ssd-81/RSS-feed-/internal/cli"
	"github.com/ssd-81/RSS-feed-/internal/config"
	"github.com/ssd-81/RSS-feed-/internal/database"
	"github.com/ssd-81/RSS-feed-/internal/types"
)

func main() {

	configData := config.Read()

	db, err := sql.Open("postgres", configData.Db_url)
	if err != nil {
		fmt.Println("database connection failed")
		os.Exit(1)
	}

	// creating new *database.Queries
	dbQueries := database.New(db)
	stateData := types.State{
		Cfg: configData,
		Db:  dbQueries,
	}

	cmds := cli.Commands{}
	// initialize map of handler functions for cmds
	cmds.Map = make(map[string]func(*types.State, cli.Command) error)
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerReset)
	cmds.Register("users", cli.HandlerUsers)
	cmds.Register("agg", cli.HandlerAgg)

	// cmds.Register("agg", rss.fetchFeed)  : this won't work, because it does not have the same function signature

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Command not found")
	}
	// creating a command struct for passing to the Run function
	cmd := cli.Command{
		Name:      args[0],
		Arguments: args[1:],
	}

	// func signature of Run : func (c *Commands) Run(s *types.State, cmd Command) error
	if err := cmds.Run(&stateData, cmd); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	} else {
		fmt.Println("command", cmd.Name, "was executed")
	}

	// fmt.Println(config.Read())

}
