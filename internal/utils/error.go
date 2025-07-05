package utils

import "github.com/breno5g/GoBudget/config"

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"error"`
}

func CheckError(err error) {
	if err != nil {
		logger := config.GetLogger("utils")
		logger.Errorf("Error: %v", err)

		panic(err)
	}
}
