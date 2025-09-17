package crawler

import (
    "context"
    "fmt"
    "sync"
)

// Shared results map
var (
    Results = make(map[string][]PageData) // store multiple pages per domain
    mu      sync.Mutex
)

func Worker(ctx context.Context, urls <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    for url := range urls {
        data, err := Crawl(ctx, url)
        if err != nil {
            fmt.Println("Failed to crawl:", url, err)
            continue
        }

        mu.Lock()
        Results[url] = append(Results[url], data) // append instead of overwrite
        mu.Unlock()
    }
}

