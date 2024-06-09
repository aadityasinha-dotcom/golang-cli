package api

import (
  "fmt"
  "github.com/spf13/cobra"
  "bytes"
  "encoding/json"
  "net/http"
)

func WeatherCommand() *cobra.Command {

  data := map[string]interface{}{
    "parent": map[string]interface{}{
      "database_id": "YOUR_DATABASE_ID",
    },
    "properties": map[string]interface{}{
      "Name": map[string]string{
        "title": "Task Name",
      },
        // Add other properties as needed
    },
  }

    // Convert data to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }

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
