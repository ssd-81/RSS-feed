package main

import (
	"fmt"

	"github.com/ssd-81/RSS-feed-/internal/config"
)

func main() {
	temp := config.Read()
	ptrTemp := config.State{}
	ptrTemp.State = &temp
	fmt.Println(ptrTemp)
	fmt.Println(ptrTemp.State.UserName)
	fmt.Println(ptrTemp.State.Db_url)
	// fmt.Println(temp)
	// temp.SetUser("alex")
	// temp = config.Read()
	// fmt.Println(temp)
}
