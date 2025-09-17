package storage

import (
    "encoding/json"
    "fmt"
    "os"
)

// SaveResults writes results to a JSON file
func SaveResults(filePath string, results interface{}) error {
    file, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("failed to create file: %v", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    if err := encoder.Encode(results); err != nil {
        return fmt.Errorf("failed to encode results: %v", err)
    }
    return nil
}
