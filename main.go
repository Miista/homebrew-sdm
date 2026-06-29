package main

import (
	"os"

	"sdm/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
