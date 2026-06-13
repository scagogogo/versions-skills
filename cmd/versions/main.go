package main

import (
	"os"

	"github.com/scagogogo/versions-skills/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
