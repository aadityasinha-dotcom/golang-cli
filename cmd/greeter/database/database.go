package database

import (
	"bytes"
	"encoding/json"
  "github.com/spf13/cobra"
	"fmt"
	"io/ioutil"
	"net/http"
)

const notionAPIVersion = "2022-06-28"

func queryDatabase(apiKey string) {
	// Create the request URL
  // https://www.notion.so/49b4ed5716ed4c089ad5c7baa32f5158?v=0ebb9632bd01425e8a5e2708f18016fe&pvs=4
  const databaseID = "49b4ed5716ed4c089ad5c7baa32f5158"
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", databaseID)

	// Create an empty JSON object for the POST request body
	jsonData := map[string]interface{}{}

	// Convert the map to JSON
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add the necessary headers
	// req.Header.Set("Authorization", "Bearer "+apiKey)
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", notionAPIVersion)

	// // Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check if the response status is not OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", body)
		return
	}

	// Print the response body
	fmt.Println("Response:", string(body))
}

func DatabaseCommand() *cobra.Command {
  return &cobra.Command{
    Use: "Api",
    Short: "API the user",
    Long: "Start the API server",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Running the API, %s!\n", args[0])
    },
  }
}
