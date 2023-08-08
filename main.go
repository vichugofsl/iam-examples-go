package main

import (
	"fmt"
	"os"

	"iam-examples-go/cmd/cli"
)

func main() {

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
