package main

import (
	"fmt"

	"github.com/breno5g/GoBudget/config"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	logger := config.GetLogger("main")
	logger.Infof("Starting server on port %s", config.GetEnv().WebServerPort)

	env := config.GetEnv()
	fmt.Println(env)
}
