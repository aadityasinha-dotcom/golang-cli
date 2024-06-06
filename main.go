package main

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
