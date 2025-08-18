package types

import ("github.com/ssd-81/RSS-feed-/internal/config"
"github.com/ssd-81/RSS-feed-/internal/database")


type State struct {
	Cfg *config.Config
	Db *database.Queries
}