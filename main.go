package main

// main function will add all the commands

import (
	"adi/cmd/greeter/root"
	"fmt"
)

func main() {
  rootCmd := root.RootCommand()

  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
  }
}
