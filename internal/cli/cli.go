package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/ssd-81/RSS-feed-/internal/database"
	"github.com/ssd-81/RSS-feed-/internal/rss"
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
			os.Exit(1)
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
		// ID:        uuid.NullUUID{UUID: uuid.NewUUID(), Valid: true},
		ID:        uuid.New(),
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

func HandlerReset(s *types.State, cmd Command) error {

	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("deletion of all rows unsuccessful")
		os.Exit(1)
	}
	err = s.Db.DeleteAllFeeds(context.Background())
	if err != nil {
		fmt.Println("deletion of all feeds unsuccessful")
		os.Exit(1)
	}
	return nil
}

func HandlerUsers(s *types.State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve users from the database")
	}
	for _, value := range users {
		if value.Name == s.Cfg.UserName {
			fmt.Println("*", value.Name, "(current)")
		} else {
			fmt.Println("*", value.Name)
		}
	}
	return nil
}

func HandlerAgg(s *types.State, cmd Command) error {
	// to be applied
	// if len(cmd.Arguments) == 0 {
	// 	return fmt.Errorf("the command agg requires a single argument: the feed url")
	// }
	url := "https://www.wagslane.dev/index.xml"
	rssData, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error encounterd : %w", err)
	}
	rss.DecodeEscapedChars(rssData)
	fmt.Println(rssData)
	fmt.Println("agg command executed successfully")

	return nil
}

func HandlerAddfeed(s *types.State, cmd Command) error {

	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("the register command expects a two argument, the name and the url")
	}

	currentUserID, err := s.Db.GetUserIDFromName(context.Background(), s.Cfg.UserName)
	if err != nil {
		fmt.Println("could not get the user id for username:", s.Cfg.UserName)
		return err
	}

	params := database.AddfeedParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      sql.NullString{String: cmd.Arguments[0], Valid: true},
		Url:       sql.NullString{String: cmd.Arguments[1], Valid: true},
		UserID:    uuid.NullUUID{UUID: currentUserID, Valid: true},
	}
	// sql.NullString{String: cmd.Arguments, Valid: true}

	feed, err := s.Db.Addfeed(context.Background(), params)
	if err != nil {
		fmt.Println("error encountered", err)
		return err
	}
	fmt.Println("feed successfully saved to database")
	fmt.Println(feed)

	return nil
}

func HandlerFeeds(s *types.State, cmd Command) error {
	// use the s.Db.GetFeeds function to get access to []Feeds, then you figure out on your own
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("their was error while retrieving feeds from the db")
	}
	for _, value := range feeds {
		fmt.Println(">", value.Name.String)
		fmt.Println(" ", value.Url.String)
		// not handling the error cases
		username, err := s.Db.GetNameFromUserID(context.Background(), value.UserID.UUID)
		if err != nil {
			fmt.Println(err, "was encountered")
		}
		fmt.Println(" ", username)
		fmt.Println()

	}
	return nil

}

func HandlerFollow(s *types.State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the command follow expects a single argument, the url of the rss feed")
	}
	// url := cmd.Arguments[0]
	// make great amount of checks in these
	currUser := s.Cfg.UserName
	userID, _ := s.Db.GetUserIDFromName(context.Background(), currUser)
	url := sql.NullString{String: cmd.Arguments[0], Valid: true}
	feedID, _ := s.Db.GetFeedIdFromUrl(context.Background(), url)

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		// UserID: userID,
		// FeedID: feedID,
		UserID: uuid.NullUUID{UUID: userID, Valid: true},
		FeedID: uuid.NullUUID{UUID: feedID, Valid: true},
	}
	feed_follow, err := s.Db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error encountered while handling command follow: %v", err)
	}
	feedName, _ := s.Db.GetFeedNameFromFeedId(context.Background(), feed_follow.ID)
	fmt.Println(feedName)
	fmt.Println(currUser)

	return nil
}

func (c *Commands) Run(s *types.State, cmd Command) error {
	// runs a given command with the provided state if it exists
	value, ok := c.Map[cmd.Name]
	if ok {
		err := value(s, cmd)
		// catching any error from the handler function
		if err != nil {
			return err
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
