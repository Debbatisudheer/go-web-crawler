package main

import (
    "context"
    "fmt"
    "sync"
    "time"

    "go-web-crawler/internal/crawler"
    "go-web-crawler/internal/storage"
)

func main() {
urls := []string{
    "https://qiskit.org/documentation/",
    "https://qiskit.org/textbook/preface.html",
    "https://qiskit.org/textbook/ch-states/single-qubit-gates.html",
    "https://qiskit.org/textbook/ch-measurement/measuring-qubits.html",
    "https://github.com/Qiskit/qiskit-tutorials",
    // add more Qiskit textbook chapters or tutorial URLs
}


    urlChan := make(chan string, len(urls))
    var wg sync.WaitGroup

    numWorkers := 3

    // Create a parent context with timeout for all workers
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go crawler.Worker(ctx, urlChan, &wg)
    }

    // Send URLs to workers
    for _, url := range urls {
        urlChan <- url
    }
    close(urlChan)

    // Wait for all workers to finish
    wg.Wait()

    // Save results to JSON file
    err := storage.SaveResults("results/results.json", crawler.Results)
    if err != nil {
        fmt.Println("Error saving results:", err)
        return
    }

    fmt.Println("Crawling completed! Results saved to results/results.json")
}
