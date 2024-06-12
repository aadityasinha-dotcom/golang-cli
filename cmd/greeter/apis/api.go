package api

import (
  "fmt"
  "github.com/spf13/cobra"
  "bytes"
  "encoding/json"
  "net/http"
)

func WeatherCommand() *cobra.Command {

    url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", "49b4ed5716ed4c089ad5c7baa32f5158")

    jsonData := map[string]interface{}{}

  client := &http.Client{}
  req, err := http.NewRequest("POST", "https://api.notion.com/v1/pages", bytes.NewBuffer(jsonData))
  if err != nil {
      panic(err)
  }

    // Include authentication token
    req.Header.Set("Authorization", "Bearer YOUR_API_KEY")
    req.Header.Set("Content-Type", "application/json")

    // Send request
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

  return &cobra.Command{
    Use: "api",
    Short: "API the user",
    Long: "Greet the user with a friendly message",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Running the API, %s!\n", args[0])
    },
  }
}
