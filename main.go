package main

import (
	"github.com/Wilddogmoto/example_project/api"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/Wilddogmoto/example_project/logging"
)

func main() {

	var (
		logger = logging.InitLogger()
		err    error
	)

	if err = data.DBConnect(); err != nil {
		logger.Fatalf("bad main connection %v\n", err)
		panic(err)
	}

	if err = api.InitRouter(); err != nil {
		logger.Fatalf("route start error %v\n", err)
		panic(err)
	}
}
