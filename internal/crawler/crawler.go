package crawler

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "github.com/PuerkitoBio/goquery"
)

// Crawl fetches a single page and extracts slokas
func Crawl(ctx context.Context, url string) (PageData, error) {
    // Retry up to 3 times
    var resp *http.Response
    var err error
    for i := 0; i < 3; i++ {
        req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
        client := &http.Client{}
        resp, err = client.Do(req)
        if err != nil {
            fmt.Printf("Retry %d for %s\n", i+1, url)
            time.Sleep(500 * time.Millisecond)
            continue
        }
        break
    }

    if err != nil {
        return PageData{}, err
    }
    defer resp.Body.Close()

    // Parse HTML with goquery
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return PageData{}, err
    }

    slokas := []string{}

    // Example: Suppose slokas are inside <p> tags
    doc.Find("p").Each(func(i int, s *goquery.Selection) {
        text := s.Text()
        if len(text) > 20 { // skip very short lines
            slokas = append(slokas, text)
        }
    })

    return PageData{
        URL:    url,
        Title:  doc.Find("title").Text(),
        Slokas: slokas,
    }, nil
}
