package main

import (
	"github.com/ulerdogan/pickaxe/cmd/pickaxe/root"
	logger "github.com/ulerdogan/pickaxe/utils/logger"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		logger.Error(err, "cannot start the app")
	}
}
