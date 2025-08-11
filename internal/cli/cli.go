package cli

import (
	"fmt"

	"github.com/ssd-81/RSS-feed-/internal/config"
)

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Map map[string]func(*config.State, Command) error
}

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	s.State.UserName = cmd.Arguments[0]
	fmt.Println("the user has been set successfully")
	return nil
}


// run  or Run
func (c *Commands) Run(s *config.State, cmd Command) error {
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
func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	// registers a new handler function for a command name
	c.Map[name] = f
	// might need to check this function again based on actual usage
}
