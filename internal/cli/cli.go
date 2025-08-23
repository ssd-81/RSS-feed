package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/ssd-81/RSS-feed-/internal/database"
	"github.com/ssd-81/RSS-feed-/internal/types"
)

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Map map[string]func(*types.State, Command) error
}

func HandlerLogin(s *types.State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	// change the logic to check if the user exists in the database;
	username := cmd.Arguments[0]
	_, err := s.Db.GetUser(context.Background(), username)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("the user does not exist in the database")
		}
		return fmt.Errorf("some unexpected error came up")

	}
	s.Cfg.SetUser(cmd.Arguments[0])
	fmt.Println("the user has been set successfully")
	return nil
}

func HandlerRegister(s *types.State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the register command expects a single argument, the username")
	}

	// checking if the user already exists in the database
	// and exit prematurely if the user already exists
	_, err := s.Db.GetUser(context.Background(), cmd.Arguments[0])

	if err == nil {
		fmt.Println("the user already exists in the database")
		os.Exit(1)
	}

	// creating arg for passing to CreateUser function
	params := database.CreateUserParams{
		ID:        uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      cmd.Arguments[0],
	}

	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("could not create the %s user in the database", params.Name)
	}
	s.Cfg.SetUser(params.Name)
	fmt.Println("user: ", user, " was successfuly registered in database")

	return nil
}

func (c *Commands) Run(s *types.State, cmd Command) error {
	// runs a given command with the provided state if it exists
	value, ok := c.Map[cmd.Name]
	if ok {
		err := value(s, cmd)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return fmt.Errorf("command does not exist in CLI")
	}
	return nil
}

func (c *Commands) Register(name string, f func(*types.State, Command) error) {
	//  registers a new handler function for a command name
	c.Map[name] = f
}
