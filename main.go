package main

import (
	"os"

	"github.com/mrxk/dev-env/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
