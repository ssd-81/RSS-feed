package main

import (
	"fmt"

	"github.com/ssd-81/RSS-feed-/internal/config"
)

func main() {
	fmt.Println("started")
	temp := config.Read()
	temp.SetUser("genie")
	newTemp := config.Read()
	fmt.Println("I am new", newTemp)
	fmt.Println(temp)
	fmt.Println("ended")
}
