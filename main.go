package main

import (
	"github.com/Wilddogmoto/example_project/api"
	"github.com/Wilddogmoto/example_project/data"
)

func main() {
	data.FindConnect()
	api.Start()
}
