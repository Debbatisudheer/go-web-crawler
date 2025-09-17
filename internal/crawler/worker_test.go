package crawler

import (
    "context"
    "sync"
    "testing"
)

func TestWorkerIntegration(t *testing.T) {
    urls := make(chan string, 1)
    wg := &sync.WaitGroup{}
    ctx := context.Background()

    urls <- "https://qiskit.org/documentation/"
    close(urls)

    wg.Add(1)
    go Worker(ctx, urls, wg)
    wg.Wait()

    if len(Results) == 0 {
        t.Error("Results should not be empty after crawling")
    }
}
