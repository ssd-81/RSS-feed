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
	s.Cfg.UserName = cmd.Arguments[0]
	fmt.Println("the user has been set successfully")
	return nil
}

func HandlerRegister(s *types.State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the register command expects a single argument, the username")
	}

	// checking if the user already exists in the database
	// and exit prematurely if the user already exists
	_, err := s.Db.GetUser(context.Background(), cmd.Arguments[1])
	if err == nil {
		os.Exit(1)
	}

	// creating arg for passing to CreateUser function
	params := database.CreateUserParams{
		ID:        uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      cmd.Arguments[1],
	}
	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("could not create the %s user in the database", params.Name)
	}
	fmt.Println("user: ", user, " was successfuly registered in database")

	return nil
}

// run  or Run
func (c *Commands) Run(s *types.State, cmd Command) error {
	// runs a given command with the provided state if it exists
	value, ok := c.Map[cmd.Name]
	if ok {
		value(s, cmd)
	} else {
		return fmt.Errorf("command does not exist in CLI")
	}
	return nil
}

// register or Register
func (c *Commands) Register(name string, f func(*types.State, Command) error) {
	// registers a new handler function for a command name
	c.Map[name] = f
	// might need to check this function again based on actual usage
}
