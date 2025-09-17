package crawler

import (
    "context"
    "testing"
)

func TestCrawl(t *testing.T) {
    url := "https://qiskit.org/documentation/"
    ctx := context.Background()

    data, err := Crawl(ctx, url)
    if err != nil {
        t.Fatalf("Crawl failed: %v", err)
    }

    if data.Title == "" {
        t.Error("Title should not be empty")
    }

    if len(data.Slokas) == 0 {
        t.Error("Content (Slokas) should not be empty")
    }
}
