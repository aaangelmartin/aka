package main

import (
	"fmt"
	"os"

	"github.com/aaangelmartin/aka/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "aka:", err)
		os.Exit(1)
	}
}
