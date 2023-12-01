package main

import (
	"fmt"
	"github.io/cbuschka/go-secret-tool/internal/cli"
	"os"
)

func main() {
	err := cli.ListSecrets()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
