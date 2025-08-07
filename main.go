package main

import (
	"fmt"

	"github.com/ssd-81/RSS-feed-/internal/config"
)

func main() {
	fmt.Println("started")
	config.Read()
}
