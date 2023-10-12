package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Define a struct to represent your JSON data structure.
type Item struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	// JSON data as a string (you can load it from a file or an HTTP request).
	jsonData := `[{"name": "Item 1", "value": 100}, {"name": "Item 2", "value": 200}, {"name": "Item 3", "value": 300}]`

	// Parse JSON data into a slice of struct.
	var items []Item
	err := json.Unmarshal([]byte(jsonData), &items)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Create a new table.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Value"})

	// Iterate through the data and add rows to the table.
	for _, item := range items {
		table.Append([]string{item.Name, fmt.Sprintf("%d", item.Value)})
	}

	// Set table properties (optional).
	table.SetCenterSeparator("│")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	table.SetBorder(false)
	table.SetAutoFormatHeaders(true)

	// Render the table.

	table.Render()
}
