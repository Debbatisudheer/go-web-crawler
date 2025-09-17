package storage

import (
    "go-web-crawler/internal/crawler"
    "os"
    "path/filepath"
    "testing"
)

func TestSaveResults(t *testing.T) {
    // Ensure results directory exists
    os.MkdirAll("results", os.ModePerm)

    tempFile := filepath.Join("results", "test_results.json")
    defer os.Remove(tempFile) // clean up after test

    testData := map[string][]crawler.PageData{
        "https://example.com": {
            {URL: "https://example.com", Title: "Example", Slokas: []string{"Test content"}},
        },
    }

    err := SaveResults(tempFile, testData)
    if err != nil {
        t.Fatalf("SaveResults failed: %v", err)
    }

    if _, err := os.Stat(tempFile); os.IsNotExist(err) {
        t.Error("File was not created")
    }
}
