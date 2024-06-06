package root

import (
  "fmt"
  "github.com/spf13/cobra"
  "adi/cmd/greeter/greet"
)

func RootCommand() *cobra.Command {
  cmd := &cobra.Command{
    Use: "greeter",
    Short: "greater is friendly",
    Long: "greeter is a friendly cli",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Welcome to Greeter")
    },
  }

  cmd.AddCommand(greet.GreetCommand())

  return cmd
}
