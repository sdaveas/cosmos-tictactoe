package main

import (
	"os"

	"github.com/sdaveas/tictactoe/cmd/tictactoed/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
