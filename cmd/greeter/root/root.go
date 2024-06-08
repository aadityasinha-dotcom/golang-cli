package root

import (
	"adi/cmd/greeter/apis"
	"adi/cmd/greeter/greet"
	"fmt"
	"github.com/spf13/cobra"
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
  cmd.AddCommand(api.WeatherCommand())

  return cmd
}
