package main

import (
	"os"

	"svctool/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
