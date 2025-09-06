package middleware

import (
	"context"
	"fmt"

	"github.com/ssd-81/RSS-feed-/internal/cli"
	"github.com/ssd-81/RSS-feed-/internal/database"
	"github.com/ssd-81/RSS-feed-/internal/types"
)

// sample usage
// cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))

func MiddlewareLoggedIn(handler func(s *types.State, cmd cli.Command, user database.User) error) func(*types.State, cli.Command) error {
	// 	func middleware(originalHandler http.Handler) http.Handler {
	//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//         fmt.Println("Running before handler")
	//         w.Write([]byte("Hijacking Request "))
	//         originalHandler.Serve(w, r)
	//         fmt.Println("Running after handler")
	//   })
	// }

	return func(state *types.State, c cli.Command) error {
		user, err := state.Db.GetUser(context.Background(), state.Cfg.UserName)
		if err != nil {
			return fmt.Errorf("error encountered while retrieving the name of current user: %v", err)
		}
		handler(state, c, user)
		return nil

	}

	// return nil

}
