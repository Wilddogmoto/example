package main

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/api"
	"github.com/Wilddogmoto/example_project/data"
)

func main() {

	if err := data.DBConnect(); err != nil {
		fmt.Println("bad main connection", err)
		return
	}

	if err := api.InitRouter(); err != nil {
		fmt.Println("route start error", err)
		return
	}
}
