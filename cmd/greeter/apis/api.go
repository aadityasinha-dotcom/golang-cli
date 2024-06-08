package api

import (
  "fmt"
  "github.com/spf13/cobra"
)

func WeatherCommand() *cobra.Command {
  return &cobra.Command{
    Use: "api",
    Short: "API the user",
    Long: "Greet the user with a friendly message",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Running the API, %s!\n", args[0])
    },
  }
}
