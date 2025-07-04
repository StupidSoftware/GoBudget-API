package utils

import "github.com/breno5g/GoBudget/config"

func CheckError(err error) {
	if err != nil {
		logger := config.GetLogger("utils")
		logger.Errorf("Error: %v", err)

		panic(err)
	}
}
