package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	err := (&cli.Command{
		Usage: "print size of a file or directory",
	}).Run(context.Background(), os.Args)
	if err != nil {
		return
	}
}
